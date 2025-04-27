package openai

type GptModel string

const (
	gpt_o4_mini   GptModel = "o4-mini"
	gpt_o3_mini   GptModel = "o3-mini"
	gpt_o1_mini   GptModel = "o1-mini"
	gpt_o3        GptModel = "o3"
	gpt_o1        GptModel = "o1"
	gpt_o1_pro    GptModel = "o1-pro"
	gpt_41        GptModel = "gpt-4.1"
	gpt_4o        GptModel = "gpt-4o"
	gpt_4o_latest GptModel = "chatgpt-4o-latest"
	gpt_41_mini   GptModel = "gpt-4.1-mini"
	gpt_4o_mini   GptModel = "gpt-4o-mini"
	gpt_41_nano   GptModel = "gpt-4.1-nano"
	gpt_35_turbo  GptModel = "gpt-3.5-turbo"
)

var GptModels = struct {
	//Reasoning models
	Gpt_o4_mini GptModel //Cost optimized
	Gpt_o3_mini GptModel //Cost optimized
	Gpt_o1_mini GptModel //Cost optimized
	Gpt_o3      GptModel
	Gpt_o1      GptModel
	Gpt_o1_pro  GptModel
	Gpt_41_mini GptModel //Cost optimized
	Gpt_4o_mini GptModel //Cost optimized
	Gpt_41_nano GptModel //Cost optimized

	//Flagship chat models
	Gpt_41        GptModel
	Gpt_4o        GptModel
	Gpt_4o_latest GptModel

	//Tool-specific models

	//Older GPT models

	//GPT base models
	Gpt_35_turbo GptModel
}{
	Gpt_o4_mini:   gpt_o4_mini,
	Gpt_o3_mini:   gpt_o3_mini,
	Gpt_o1_mini:   gpt_o1_mini,
	Gpt_o3:        gpt_o3,
	Gpt_o1:        gpt_o1,
	Gpt_o1_pro:    gpt_o1_pro,
	Gpt_41:        gpt_41,
	Gpt_4o:        gpt_4o,
	Gpt_4o_latest: gpt_4o_latest,
	Gpt_41_mini:   gpt_41_mini,
	Gpt_4o_mini:   gpt_4o_mini,
	Gpt_41_nano:   gpt_41_nano,
	Gpt_35_turbo:  gpt_35_turbo,
}
