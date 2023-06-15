package aigcaas

// CommonText2ImgRequest  文生图通用的请求参数
type CommonText2ImgRequest struct {
	EnableHr                          bool          `json:"enable_hr,omitempty"`
	DenoisingStrength                 float64       `json:"denoising_strength,omitempty"`
	FirstphaseWidth                   int           `json:"firstphase_width,omitempty"`
	FirstphaseHeight                  int           `json:"firstphase_height,omitempty"`
	HrScale                           int           `json:"hr_scale,omitempty"`
	HrUpscaler                        string        `json:"hr_upscaler,omitempty"`
	HrSecondPassSteps                 int           `json:"hr_second_pass_steps,omitempty"`
	HrResizeX                         int           `json:"hr_resize_x,omitempty"`
	HrResizeY                         int           `json:"hr_resize_y,omitempty"`
	HrSamplerName                     string        `json:"hr_sampler_name,omitempty"`
	HrPrompt                          string        `json:"hr_prompt,omitempty"`
	HrNegativePrompt                  string        `json:"hr_negative_prompt,omitempty"`
	Prompt                            string        `json:"prompt,omitempty"`
	Styles                            []string      `json:"styles,omitempty"`
	Seed                              int           `json:"seed,omitempty"`
	Subseed                           int           `json:"subseed,omitempty"`
	SubseedStrength                   int           `json:"subseed_strength,omitempty"`
	SeedResizeFromH                   int           `json:"seed_resize_from_h,omitempty"`
	SeedResizeFromW                   int           `json:"seed_resize_from_w,omitempty"`
	SamplerName                       string        `json:"sampler_name,omitempty"`
	BatchSize                         int           `json:"batch_size,omitempty"`
	NIter                             int           `json:"n_iter,omitempty"`
	Steps                             int           `json:"steps,omitempty"`
	CfgScale                          int           `json:"cfg_scale,omitempty"`
	Width                             int           `json:"width,omitempty"`
	Height                            int           `json:"height,omitempty"`
	RestoreFaces                      bool          `json:"restore_faces,omitempty"`
	Tiling                            bool          `json:"tiling,omitempty"`
	DoNotSaveSamples                  bool          `json:"do_not_save_samples,omitempty"`
	DoNotSaveGrid                     bool          `json:"do_not_save_grid,omitempty"`
	NegativePrompt                    string        `json:"negative_prompt,omitempty"`
	Eta                               int           `json:"eta,omitempty"`
	SMinUncond                        int           `json:"s_min_uncond,omitempty"`
	SChurn                            int           `json:"s_churn,omitempty"`
	STmax                             int           `json:"s_tmax,omitempty"`
	STmin                             int           `json:"s_tmin,omitempty"`
	SNoise                            int           `json:"s_noise,omitempty"`
	OverrideSettings                  *struct{}     `json:"override_settings,omitempty"`
	OverrideSettingsRestoreAfterwards bool          `json:"override_settings_restore_afterwards,omitempty"`
	ScriptArgs                        []interface{} `json:"script_args,omitempty"`
	SamplerIndex                      string        `json:"sampler_index,omitempty"`
	ScriptName                        string        `json:"script_name,omitempty"`
	SendImages                        bool          `json:"send_images,omitempty"`
	SaveImages                        bool          `json:"save_images,omitempty"`
	AlwaysonScripts                   *struct{}     `json:"alwayson_scripts,omitempty"`
}

// CommonImg2ImgRequest 通用图生图请求参数
type CommonImg2ImgRequest struct {
	InitImages                        []string      `json:"init_images,omitempty"`
	ResizeMode                        int           `json:"resize_mode,omitempty"`
	DenoisingStrength                 float64       `json:"denoising_strength,omitempty"`
	ImageCfgScale                     int           `json:"image_cfg_scale,omitempty"`
	Mask                              string        `json:"mask,omitempty"`
	MaskBlur                          int           `json:"mask_blur,omitempty"`
	InpaintingFill                    int           `json:"inpainting_fill,omitempty"`
	InpaintFullRes                    bool          `json:"inpaint_full_res,omitempty"`
	InpaintFullResPadding             int           `json:"inpaint_full_res_padding,omitempty"`
	InpaintingMaskInvert              int           `json:"inpainting_mask_invert,omitempty"`
	InitialNoiseMultiplier            int           `json:"initial_noise_multiplier,omitempty"`
	Prompt                            string        `json:"prompt,omitempty"`
	Styles                            []string      `json:"styles,omitempty"`
	Seed                              int           `json:"seed,omitempty"`
	Subseed                           int           `json:"subseed,omitempty"`
	SubseedStrength                   int           `json:"subseed_strength,omitempty"`
	SeedResizeFromH                   int           `json:"seed_resize_from_h,omitempty"`
	SeedResizeFromW                   int           `json:"seed_resize_from_w,omitempty"`
	SamplerName                       string        `json:"sampler_name,omitempty"`
	BatchSize                         int           `json:"batch_size,omitempty"`
	NIter                             int           `json:"n_iter,omitempty"`
	Steps                             int           `json:"steps,omitempty"`
	CfgScale                          int           `json:"cfg_scale,omitempty"`
	Width                             int           `json:"width,omitempty"`
	Height                            int           `json:"height,omitempty"`
	RestoreFaces                      bool          `json:"restore_faces,omitempty"`
	Tiling                            bool          `json:"tiling,omitempty"`
	DoNotSaveSamples                  bool          `json:"do_not_save_samples,omitempty"`
	DoNotSaveGrid                     bool          `json:"do_not_save_grid,omitempty"`
	NegativePrompt                    string        `json:"negative_prompt,omitempty"`
	Eta                               int           `json:"eta,omitempty"`
	SMinUncond                        int           `json:"s_min_uncond,omitempty"`
	SChurn                            int           `json:"s_churn,omitempty"`
	STmax                             int           `json:"s_tmax,omitempty"`
	STmin                             int           `json:"s_tmin,omitempty"`
	SNoise                            int           `json:"s_noise,omitempty"`
	OverrideSettings                  *struct{}     `json:"override_settings,omitempty"`
	OverrideSettingsRestoreAfterwards bool          `json:"override_settings_restore_afterwards,omitempty"`
	ScriptArgs                        []interface{} `json:"script_args,omitempty"`
	SamplerIndex                      string        `json:"sampler_index,omitempty"`
	IncludeInitImages                 bool          `json:"include_init_images,omitempty"`
	ScriptName                        string        `json:"script_name,omitempty"`
	SendImages                        bool          `json:"send_images,omitempty"`
	SaveImages                        bool          `json:"save_images,omitempty"`
	AlwaysonScripts                   *struct{}     `json:"alwayson_scripts,omitempty"`
}

// CommonSuccessResponse 通用的请求成功的响应参数
type CommonSuccessResponse struct {
	AigcaasRequestId string `json:"Aigcaas-Request-Id"`
}

// CommonErrorResponse 通用的请求失败的响应参数
type CommonErrorResponse struct {
	Status  string `json:"status"`
	Message string `json:"message"`
	Help    string `json:"help"`
}

// AsyncRequestIdErrorResponse 失败异步调用
type AsyncRequestIdErrorResponse struct {
	Error   string `json:"error"`
	Message string `json:"message"`
	Help    string `json:"help"`
}

// AsyncRequestIdSuccessResponse 成功异步调用
type AsyncRequestIdSuccessResponse struct {
	Images     []string `json:"images"`
	Parameters struct {
		EnableHr          bool        `json:"enable_hr"`
		DenoisingStrength float64     `json:"denoising_strength"`
		FirstphaseWidth   int         `json:"firstphase_width"`
		FirstphaseHeight  int         `json:"firstphase_height"`
		HrScale           float64     `json:"hr_scale"`
		HrUpscaler        interface{} `json:"hr_upscaler"`
		HrSecondPassSteps int         `json:"hr_second_pass_steps"`
		HrResizeX         int         `json:"hr_resize_x"`
		HrResizeY         int         `json:"hr_resize_y"`
		HrSamplerName     interface{} `json:"hr_sampler_name"`
		HrPrompt          string      `json:"hr_prompt"`
		HrNegativePrompt  string      `json:"hr_negative_prompt"`
		Prompt            string      `json:"prompt"`
		Styles            interface{} `json:"styles"`
		Seed              int         `json:"seed"`
		Subseed           int         `json:"subseed"`
		SubseedStrength   int         `json:"subseed_strength"`
		SeedResizeFromH   int         `json:"seed_resize_from_h"`
		SeedResizeFromW   int         `json:"seed_resize_from_w"`
		SamplerName       interface{} `json:"sampler_name"`
		BatchSize         int         `json:"batch_size"`
		NIter             int         `json:"n_iter"`
		Steps             int         `json:"steps"`
		CfgScale          float64     `json:"cfg_scale"`
		Width             int         `json:"width"`
		Height            int         `json:"height"`
		RestoreFaces      bool        `json:"restore_faces"`
		Tiling            bool        `json:"tiling"`
		DoNotSaveSamples  bool        `json:"do_not_save_samples"`
		DoNotSaveGrid     bool        `json:"do_not_save_grid"`
		NegativePrompt    interface{} `json:"negative_prompt"`
		Eta               interface{} `json:"eta"`
		SMinUncond        float64     `json:"s_min_uncond"`
		SChurn            float64     `json:"s_churn"`
		STmax             interface{} `json:"s_tmax"`
		STmin             float64     `json:"s_tmin"`
		SNoise            float64     `json:"s_noise"`
		OverrideSettings  struct {
		} `json:"override_settings"`
		OverrideSettingsRestoreAfterwards bool          `json:"override_settings_restore_afterwards"`
		ScriptArgs                        []interface{} `json:"script_args"`
		SamplerIndex                      string        `json:"sampler_index"`
		ScriptName                        interface{}   `json:"script_name"`
		SendImages                        bool          `json:"send_images"`
		SaveImages                        bool          `json:"save_images"`
		AlwaysonScripts                   struct {
		} `json:"alwayson_scripts"`
	} `json:"parameters"`
	Info string `json:"info"`
}

// AsyncRequestIdInfoSuccessResponse 成功异步调用中info响应参数
type AsyncRequestIdInfoSuccessResponse struct {
	Prompt                        string        `json:"prompt"`
	AllPrompts                    []string      `json:"all_prompts"`
	NegativePrompt                string        `json:"negative_prompt"`
	AllNegativePrompts            []string      `json:"all_negative_prompts"`
	Seed                          int           `json:"seed"`
	AllSeeds                      []int         `json:"all_seeds"`
	Subseed                       int           `json:"subseed"`
	AllSubseeds                   []int         `json:"all_subseeds"`
	SubseedStrength               int           `json:"subseed_strength"`
	Width                         int           `json:"width"`
	Height                        int           `json:"height"`
	SamplerName                   string        `json:"sampler_name"`
	CfgScale                      float64       `json:"cfg_scale"`
	Steps                         int           `json:"steps"`
	BatchSize                     int           `json:"batch_size"`
	RestoreFaces                  bool          `json:"restore_faces"`
	FaceRestorationModel          interface{}   `json:"face_restoration_model"`
	SdModelHash                   string        `json:"sd_model_hash"`
	SeedResizeFromW               int           `json:"seed_resize_from_w"`
	SeedResizeFromH               int           `json:"seed_resize_from_h"`
	DenoisingStrength             int           `json:"denoising_strength"`
	ExtraGenerationParams         struct{}      `json:"extra_generation_params"`
	IndexOfFirstImage             int           `json:"index_of_first_image"`
	Infotexts                     []string      `json:"infotexts"`
	Styles                        []interface{} `json:"styles"`
	JobTimestamp                  string        `json:"job_timestamp"`
	ClipSkip                      int           `json:"clip_skip"`
	IsUsingInpaintingConditioning bool          `json:"is_using_inpainting_conditioning"`
}

// AsyncRequestIdWaitResponse 成功异步调用 等待结果
type AsyncRequestIdWaitResponse struct {
	Message string `json:"message"`
	Help    string `json:"help"`
}

// ComputerVisionRequest 计算机视觉分类模型请求参数
type ComputerVisionRequest struct {
	ImageUrl    string `json:"image_url,omitempty"`
	ImageBase64 string `json:"image_base64,omitempty"`
}

// ComputerVisionSuccessResponse 关于计算机视觉分类成功的响应参数
type ComputerVisionSuccessResponse struct {
	Status string `json:"status"`
	Data   string `json:"data"`
}

// ComputerVisionErrorResponse 关于计算机视觉分类失败的响应参数
type ComputerVisionErrorResponse struct {
	Error   string `json:"error"`
	Message struct {
		Error   string `json:"error"`
		Message string `json:"message"`
	} `json:"message"`
	Help string `json:"help"`
}

// GeneralRecognitionSuccessResponse 万物识别成功的响应参数
type GeneralRecognitionSuccessResponse struct {
	Status string `json:"status"`
	Data   struct {
		Scores []float64 `json:"scores"`
		Labels []string  `json:"labels"`
	} `json:"data"`
}
