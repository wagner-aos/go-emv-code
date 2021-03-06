package mpm_test

import (
	"testing"

	"go.mercari.io/go-emv-code/mpm"
)

func TestTipOrConvenienceIndicator_Tokenize(t *testing.T) {
	tests := []struct {
		name    string
		give    *mpm.TipOrConvenienceIndicator
		want    string
		wantErr bool
	}{
		{
			name:    "give nil",
			give:    nil,
			want:    "",
			wantErr: false,
		},
		{
			name: "give mpm.TipOrConvenienceIndicatorPrompt",
			give: func() *mpm.TipOrConvenienceIndicator {
				p := mpm.TipOrConvenienceIndicatorPrompt
				return &p
			}(),
			want:    "01",
			wantErr: false,
		},
		{
			name: "give mpm.TipOrConvenienceIndicatorFixed",
			give: func() *mpm.TipOrConvenienceIndicator {
				p := mpm.TipOrConvenienceIndicatorFixed
				return &p
			}(),
			want:    "02",
			wantErr: false,
		},
		{
			name: "give mpm.TipOrConvenienceIndicatorPercentage",
			give: func() *mpm.TipOrConvenienceIndicator {
				p := mpm.TipOrConvenienceIndicatorPercentage
				return &p
			}(),
			want:    "03",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			dst, err := tt.give.Tokenize()

			if (err != nil) != tt.wantErr {
				t.Errorf("TipOrConvenienceIndicator.Tokenize error = %v, wantErr %v", err, tt.wantErr)
			}

			if tt.want != dst {
				t.Errorf("TipOrConvenienceIndicator.Tokenize = %v, want %v", dst, tt.want)
			}
		})
	}
}

func TestTipOrConvenienceIndicator_Scan(t *testing.T) {
	tests := []struct {
		name    string
		give    []rune
		wantErr bool
	}{
		{
			name:    "give nil",
			give:    nil,
			wantErr: true,
		},
		{
			name:    "give empty",
			give:    []rune{},
			wantErr: true,
		},
		{
			name:    "give unexpected string",
			give:    []rune("wrong_value"),
			wantErr: true,
		},
		{
			name:    "give mpm.TipOrConvenienceIndicatorPrompt",
			give:    []rune(mpm.TipOrConvenienceIndicatorPrompt),
			wantErr: false,
		},
		{
			name:    "give mpm.TipOrConvenienceIndicatorFixed",
			give:    []rune(mpm.TipOrConvenienceIndicatorFixed),
			wantErr: false,
		},
		{
			name:    "give mpm.TipOrConvenienceIndicatorPercentage",
			give:    []rune(mpm.TipOrConvenienceIndicatorPercentage),
			wantErr: false,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			var ind mpm.TipOrConvenienceIndicator
			err := ind.Scan(tt.give)

			if (err != nil) != tt.wantErr {
				t.Errorf("TipOrConvenienceIndicator.Scan error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestPointOfInitiationMethod_Tokenize(t *testing.T) {
	tests := []struct {
		name    string
		give    *mpm.PointOfInitiationMethod
		want    string
		wantErr bool
	}{
		{
			name:    "give nil",
			give:    nil,
			want:    "",
			wantErr: false,
		},
		{
			name: "give mpm.PointOfInitiationMethodStatic",
			give: func() *mpm.PointOfInitiationMethod {
				p := mpm.PointOfInitiationMethodStatic
				return &p
			}(),
			want:    "11",
			wantErr: false,
		},
		{
			name: "give mpm.PointOfInitiationMethodDynamic",
			give: func() *mpm.PointOfInitiationMethod {
				p := mpm.PointOfInitiationMethodDynamic
				return &p
			}(),
			want:    "12",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			dst, err := tt.give.Tokenize()

			if (err != nil) != tt.wantErr {
				t.Errorf("PointOfInitiationMethod.Tokenize error = %v, wantErr %v", err, tt.wantErr)
			}

			if tt.want != dst {
				t.Errorf("PointOfInitiationMethod.Tokenize = %v, want %v", dst, tt.want)
			}
		})
	}
}

func TestPointOfInitiationMethod_Scan(t *testing.T) {
	tests := []struct {
		name    string
		give    []rune
		wantErr bool
	}{
		{
			name:    "give nil",
			give:    nil,
			wantErr: true,
		},
		{
			name:    "give empty",
			give:    []rune{},
			wantErr: true,
		},
		{
			name:    "give unexpected string",
			give:    []rune("wrong_value"),
			wantErr: true,
		},
		{
			name:    "give mpm.PointOfInitiationMethodStatic",
			give:    []rune(mpm.PointOfInitiationMethodStatic),
			wantErr: false,
		},
		{
			name:    "give mpm.PointOfInitiationMethodDynamic",
			give:    []rune(mpm.PointOfInitiationMethodDynamic),
			wantErr: false,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			var ind mpm.PointOfInitiationMethod
			err := ind.Scan(tt.give)

			if (err != nil) != tt.wantErr {
				t.Errorf("PointOfInitiationMethod.Scan error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestNullString_Tokenize(t *testing.T) {
	tests := []struct {
		name    string
		give    *mpm.NullString
		want    string
		wantErr bool
	}{
		{
			name:    "give nil",
			give:    nil,
			want:    "",
			wantErr: false,
		},
		{
			name:    "give invalid mpm.NullString",
			give:    &mpm.NullString{},
			want:    "",
			wantErr: false,
		},
		{
			name: "give invalid mpm.NullString",
			give: &mpm.NullString{
				String: "good_value",
				Valid:  true,
			},
			want:    "good_value",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			dst, err := tt.give.Tokenize()

			if (err != nil) != tt.wantErr {
				t.Errorf("NullString.Tokenize error = %v, wantErr %v", err, tt.wantErr)
			}

			if tt.want != dst {
				t.Errorf("NullString.Tokenize = %v, want %v", dst, tt.want)
			}
		})
	}
}

func TestNullString_Scan(t *testing.T) {
	tests := []struct {
		name    string
		give    []rune
		wantErr bool
	}{
		{
			name:    "give nil",
			give:    nil,
			wantErr: false,
		},
		{
			name:    "give empty",
			give:    []rune{},
			wantErr: false,
		},
		{
			name:    "give string",
			give:    []rune("good_value"),
			wantErr: false,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			var ind mpm.NullString
			err := ind.Scan(tt.give)

			if (err != nil) != tt.wantErr {
				t.Errorf("NullString.Scan error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestNullMerchantInformation_Tokenize(t *testing.T) {
	tests := []struct {
		name    string
		give    *mpm.NullMerchantInformation
		want    string
		wantErr bool
	}{
		{
			name:    "give nil",
			give:    nil,
			want:    "",
			wantErr: false,
		},
		{
			name:    "give invalid mpm.NullString",
			give:    &mpm.NullMerchantInformation{},
			want:    "",
			wantErr: false,
		},
		{
			name: "give empty mpm.NullString",
			give: &mpm.NullMerchantInformation{
				Valid: true,
			},
			want:    "",
			wantErr: false,
		},
		{
			name: "give valid mpm.NullString",
			give: &mpm.NullMerchantInformation{
				LanguagePreference: "ZH",
				Name:               "最佳运输",
				City:               "北京",
				Valid:              true,
			},
			want:    "0002ZH0104最佳运输0202北京",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			dst, err := tt.give.Tokenize()

			if (err != nil) != tt.wantErr {
				t.Errorf("NullMerchantInformation.Tokenize error = %v, wantErr %v", err, tt.wantErr)
			}

			if tt.want != dst {
				t.Errorf("NullMerchantInformation.Tokenize = %v, want %v", dst, tt.want)
			}
		})
	}
}

func TestNullMerchantInformation_Scan(t *testing.T) {
	tests := []struct {
		name    string
		give    []rune
		wantErr bool
	}{
		{
			name:    "give nil",
			give:    nil,
			wantErr: false,
		},
		{
			name:    "give empty",
			give:    []rune{},
			wantErr: false,
		}, {
			name:    "give wrong string",
			give:    []rune("wrong_value"),
			wantErr: true,
		},
		{
			name:    "give string",
			give:    []rune("0202北京"),
			wantErr: false,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			var ind mpm.NullMerchantInformation
			err := ind.Scan(tt.give)

			if (err != nil) != tt.wantErr {
				t.Errorf("NullMerchantInformation.Scan error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
