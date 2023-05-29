package listeners

import (
	"fmt"
	"strings"

	"github.com/antlr4-go/antlr/v4"
	"github.com/saqura-io/hanami/petalscript"
	"github.com/saqura-io/hanami/util"
)

type petalListener struct {
	petalscript.BasePetalScriptListener
	stack  []Value
	Result string
}

func NewPetalListener() *petalListener {
	return &petalListener{
		stack: []Value{},
	}
}

func (l *petalListener) push(v Value) {
	l.stack = append(l.stack, v)
}

func (l *petalListener) pop() Value {
	if len(l.stack) < 1 {
		panic("pop from empty stack")
	}

	// Get the last value from the stack
	result := l.stack[len(l.stack)-1]

	// Remove the last element from the stack
	l.stack = l.stack[:len(l.stack)-1]

	return result
}

func (l *petalListener) EnterEveryRule(ctx antlr.ParserRuleContext) {
	// fmt.Println(ctx.GetText())
}

type Value interface {
	String() string
}

type StringValue struct {
	value string
}

func (sv *StringValue) String() string {
	return sv.value
}

type FilterCall struct {
	from   *StringValue
	by     *StringValue
	equals *StringValue
}

func (fc *FilterCall) String() string {
	// Convert filter call into SQL
	return fmt.Sprintf("SELECT * FROM logs WHERE application_id = %s AND %s = %s", fc.from.value, util.StripQuotes(fc.by.value), fc.equals.value)
}

type Aggregate struct {
	calls []Value
}

func (ag *Aggregate) String() string {
	result := make([]string, len(ag.calls))

	for i, call := range ag.calls {
		result[i] = call.String()
	}

	return strings.Join(result, "\nUNION\n")
}

type OutputStatement struct {
	format *StringValue
	value  Value
}

func (os *OutputStatement) String() string {
	// Only returns spring representation of the value for simplicity
	// TODO: Actually do something useful with it
	return os.value.String()
}

func (l *petalListener) EnterFilter_call(ctx *petalscript.Filter_callContext) {
	call := &FilterCall{
		from: &StringValue{
			value: ctx.STRING(0).GetText(),
		},
		by: &StringValue{
			value: ctx.STRING(1).GetText(),
		},
		equals: &StringValue{
			value: ctx.STRING(2).GetText(),
		},
	}
	l.push(call)
}

func (l *petalListener) ExitFilter_call(ctx *petalscript.Filter_callContext) {
	if len(l.stack) == 0 {
		fmt.Println("Error: Attempt to pop from empty stack")
		return
	}

	call, ok := l.pop().(*FilterCall)
	if !ok {
		fmt.Println("Error: Failed to pop FilterCall from stack")
	}

	l.stack = append(l.stack, call)
}

func (l *petalListener) EnterAggregate(ctx *petalscript.AggregateContext) {
	aggregate := &Aggregate{}
	l.push(aggregate)

}

func (l *petalListener) ExitAggregate(ctx *petalscript.AggregateContext) {
	aggregate := l.pop().(*Aggregate)

	for range ctx.AllIdentifier() {
		aggregate.calls = append(aggregate.calls, l.pop())
	}

	l.Result = aggregate.String()
}

func (l *petalListener) EnterOutput_statement(ctx *petalscript.Output_statementContext) {
	outputStatement := &OutputStatement{}
	l.push(outputStatement)
}

func (l *petalListener) ExitOutput_statement(ctx *petalscript.Output_statementContext) {
	outputStatement := l.pop().(*OutputStatement)
	outputStatement.value = l.pop().(Value)
	outputStatement.format = l.pop().(*StringValue)
}
