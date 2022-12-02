package cast_anything

// ------------------------------------------------ ---------------------------------------------------------------------

type ValueCastRule int

const (

	// ToDefaultZeroValue 转换为对应类型默认的零值
	ToDefaultZeroValue ValueCastRule = iota

	// ToForceCast 做一个强制类型转换
	ToForceCast

	// ToError 报错，认为是处理不了
	ToError
)

// ------------------------------------------------ Nil ----------------------------------------------------------------

// NilCastRule nil值应该如何转换
type NilCastRule = int

const (

	// NilToZeroValue 转换为对应类型的零值
	NilToZeroValue NilCastRule = iota

	// NilToError 直接返回错误
	NilToError
)

// ------------------------------------------------ NegativeSignedToUnsignedCastRule ---------------------------------------

type NegativeSignedToUnsignedCastRule int

const (

	// NegativeSignedToUnsignedRuleToZero 转换为零值
	NegativeSignedToUnsignedRuleToZero NegativeSignedToUnsignedCastRule = iota

	// NegativeSignedToUnsignedRuleAbs 转换为绝对值
	NegativeSignedToUnsignedRuleAbs

	// NegativeSignedToUnsignedRuleError 转换为错误
	NegativeSignedToUnsignedRuleError
)

// ------------------------------------------------ ---------------------------------------------------------------------

type Options struct {

	// 如果被转换的类型是float类型时如何处理
	Float ValueCastRule

	// 如果被转换的类型是空字符串时如何处理
	EmptyString ValueCastRule

	// 如果被转换的类型是nil时如何处理
	Nil NilCastRule

	// 有符号类型的负数向无符号类型转换时应该如何处理
	NegativeSignedToUnsigned NegativeSignedToUnsignedCastRule
}

// ------------------------------------------------ ---------------------------------------------------------------------
