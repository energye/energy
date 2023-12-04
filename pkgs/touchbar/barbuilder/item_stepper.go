package barbuilder

type StepperFormat string

const (
	StepperNone       StepperFormat = "none"
	StepperDecimal    StepperFormat = "decimal"
	StepperPercent    StepperFormat = "percent"
	StepperScientific StepperFormat = "scientific"
	StepperOrdinal    StepperFormat = "ordinal"
	StepperCurrency   StepperFormat = "currency"
	StepperSpellOut   StepperFormat = "spell_out"
	StepperBytes      StepperFormat = "bytes"
	StepperDate       StepperFormat = "date"
	StepperDuration   StepperFormat = "duration"
)

type StepperOnChange func(value float64)

type Stepper struct {
	CommonProperties

	Value     float64
	Minimum   float64
	Maximum   float64
	Increment float64
	Format    StepperFormat
	OnChange  StepperOnChange
}

var _ Item = &Stepper{}

func (me *Stepper) isAnItem() {}
