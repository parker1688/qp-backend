package tool

import (
	"github.com/shopspring/decimal"
)

func Decimal(value float64) float64 {
	return decimal.NewFromFloat(value).Truncate(2).InexactFloat64()
}

func DecimalBit(value float64, b int) float64 {
	return decimal.NewFromFloat(value).Truncate(int32(b)).InexactFloat64()
}

// 浮点数相加，prec 保留几位小数
func DecimalFAddTruncate(a, b float64, prec int) float64 {
	aD := decimal.NewFromFloat(a)
	tmpPrec := int32(prec)

	result := aD.Add(decimal.NewFromFloat(b)).Truncate(tmpPrec).InexactFloat64()

	return result
}

// 浮点数相减，prec 保留几位小数
func DecimalFSubTruncate(a, b float64, prec int) float64 {
	aD := decimal.NewFromFloat(a)
	tmpPrec := int32(prec)

	result := aD.Sub(decimal.NewFromFloat(b)).Truncate(tmpPrec).InexactFloat64()

	return result
}

// 浮点数相乘，prec 保留几位小数
func DecimalFMulTruncate(a, b float64, prec int) float64 {
	aD := decimal.NewFromFloat(a)
	tmpPrec := int32(prec)

	result := aD.Mul(decimal.NewFromFloat(b)).Truncate(tmpPrec).InexactFloat64()

	return result
}

// 浮点数相乘，prec 保留几位小数
func DecimalFDivTruncate(a, b float64, prec int) float64 {
	if b == 0.00 {
		return 0.00
	}

	aD := decimal.NewFromFloat(a)
	tmpPrec := int32(prec)

	result := aD.Div(decimal.NewFromFloat(b)).Truncate(tmpPrec).InexactFloat64()

	return result
}

// 浮点数相加，prec 保留几位小数
func DecimalFAddRound(a, b float64, prec int) float64 {
	aD := decimal.NewFromFloat(a)
	tmpPrec := int32(prec)

	result := aD.Add(decimal.NewFromFloat(b)).Round(tmpPrec).InexactFloat64()

	return result
}

// 浮点数相减，prec 保留几位小数
func DecimalFSubRound(a, b float64, prec int) float64 {
	aD := decimal.NewFromFloat(a)
	tmpPrec := int32(prec)

	result := aD.Sub(decimal.NewFromFloat(b)).Round(tmpPrec).InexactFloat64()

	return result
}

// 浮点数相乘，prec 保留几位小数
func DecimalFMulRound(a, b float64, prec int) float64 {
	aD := decimal.NewFromFloat(a)
	tmpPrec := int32(prec)

	result := aD.Mul(decimal.NewFromFloat(b)).Round(tmpPrec).InexactFloat64()

	return result
}

// 浮点数相乘，prec 保留几位小数
func DecimalFDivRound(a, b float64, prec int) float64 {
	if b == 0.00 {
		return 0.00
	}

	aD := decimal.NewFromFloat(a)
	tmpPrec := int32(prec)

	result := aD.Div(decimal.NewFromFloat(b)).Round(tmpPrec).InexactFloat64()

	return result
}

// 浮点数相加，prec 保留几位小数
func DecimalIAddTruncate(a, b int64, prec int) float64 {
	aD := decimal.NewFromInt(a)
	tmpPrec := int32(prec)

	result := aD.Add(decimal.NewFromInt(b)).Truncate(tmpPrec).InexactFloat64()

	return result
}

// 浮点数相减，prec 保留几位小数
func DecimalISubTruncate(a, b int64, prec int) float64 {
	aD := decimal.NewFromInt(a)
	tmpPrec := int32(prec)

	result := aD.Sub(decimal.NewFromInt(b)).Truncate(tmpPrec).InexactFloat64()

	return result
}

// 浮点数相乘，prec 保留几位小数
func DecimalIMulTruncate(a, b int64, prec int) float64 {
	aD := decimal.NewFromInt(a)
	tmpPrec := int32(prec)

	result := aD.Mul(decimal.NewFromInt(b)).Truncate(tmpPrec).InexactFloat64()

	return result
}

// 浮点数相乘，prec 保留几位小数
func DecimalIDivTruncate(a, b int64, prec int) float64 {
	if b == 0.00 {
		return 0.00
	}

	aD := decimal.NewFromInt(a)
	tmpPrec := int32(prec)

	result := aD.Div(decimal.NewFromInt(b)).Truncate(tmpPrec).InexactFloat64()

	return result
}

// 浮点数相加，prec 保留几位小数
func DecimalIAddRound(a, b int64, prec int) float64 {
	aD := decimal.NewFromInt(a)
	tmpPrec := int32(prec)

	result := aD.Add(decimal.NewFromInt(b)).Round(tmpPrec).InexactFloat64()

	return result
}

// 浮点数相减，prec 保留几位小数
func DecimalISubRound(a, b int64, prec int) float64 {
	aD := decimal.NewFromInt(a)
	tmpPrec := int32(prec)

	result := aD.Sub(decimal.NewFromInt(b)).Round(tmpPrec).InexactFloat64()

	return result
}

// 浮点数相乘，prec 保留几位小数
func DecimalIMulRound(a, b int64, prec int) float64 {
	aD := decimal.NewFromInt(a)
	tmpPrec := int32(prec)

	result := aD.Mul(decimal.NewFromInt(b)).Round(tmpPrec).InexactFloat64()

	return result
}

// 浮点数相乘，prec 保留几位小数
func DecimalIDivRound(a, b int64, prec int) float64 {
	if b == 0.00 {
		return 0.00
	}

	aD := decimal.NewFromInt(a)
	tmpPrec := int32(prec)

	result := aD.Div(decimal.NewFromInt(b)).Round(tmpPrec).InexactFloat64()

	return result
}
