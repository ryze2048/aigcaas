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
