package aigcaas

// 基于Stable Diffusion模型
const (
	Analog                                    string = `analog`                                     // 胶片质感扩散生成模型 https://www.aigcaas.cn/product/detail/26.html
	Modi                                      string = `modi`                                       // 迪士尼风格扩散生成模型 https://www.aigcaas.cn/product/detail/27.html
	Sd15                                      string = `sd15`                                       // Stable Diffusion 1.5 https://www.aigcaas.cn/product/detail/50.html
	Sd21                                      string = `sd21`                                       // Stable Diffusion 2.1 https://www.aigcaas.cn/product/detail/36.html
	Openjourney                               string = `openjourney`                                // Open Journey https://www.aigcaas.cn/product/detail/66.html
	SciFiDiffusionV10V10                      string = `sciFiDiffusionV10_v10`                      // Sci-Fi Diffusion v1.0 https://www.aigcaas.cn/product/detail/195.html
	DiffusionBrushEverythingSFWNSFWAllV10     string = `diffusionBrushEverythingSFWNSFWAll_v10`     // Diffusion Brush ( Everything ) - SFW / NSFW- All Purpose Checkpoint - [ Nuclear Diffusion ] - [ Anime Hybrid ] https://www.aigcaas.cn/product/detail/194.html
	ReliberateV10                             string = `reliberate_v10`                             // Reliberate https://www.aigcaas.cn/product/detail/193.html
	GhostmixV20Bakedvae                       string = `ghostmix_v20Bakedvae`                       // GhostMix https://www.aigcaas.cn/product/detail/192.html
	PastelMixStylizedAnimePastelMixPrunedFP16 string = `pastelMixStylizedAnime_pastelMixPrunedFP16` // Pastel-Mix [Stylized Anime Model] https://www.aigcaas.cn/product/detail/191.html
	CyberrealisticV31                         string = `cyberrealistic_v31`                         // CyberRealistic https://www.aigcaas.cn/product/detail/190.html
	MajicmixRealisticV5                       string = `majicmixRealistic_v5`                       // majicMIX realistic https://www.aigcaas.cn/product/detail/189.html
	AnythingV5V5PrtRE                         string = `AnythingV5_v5PrtRE`                         // 万象熔炉 | Anything V5/Ink https://www.aigcaas.cn/product/detail/188.html
	Babes20                                   string = `babes_20`                                   // Babes https://www.aigcaas.cn/product/detail/187.html
	Level4V50BakedVAEFp16                     string = `level4_v50BakedVAEFp16`                     // Level4 https://www.aigcaas.cn/product/detail/186.html
	NeonIsometric1                            string = `neonIsometric_1`                            // Neon Isometric https://www.aigcaas.cn/product/detail/185.html
	MixrealV2Sd21MixrealV2                    string = `mixrealV2Sd21_mixrealV2`                    // MixReal v2 _ Sd 2.1 https://www.aigcaas.cn/product/detail/184.html
	FruitFusionVer1OFR2                       string = `fruitFusion_ver1OFR2`                       // Fruit Fusion https://www.aigcaas.cn/product/detail/183.html
	IsometricDreamsSD211                      string = `isometricDreamsSD21_1`                      // Isometric Dreams SD2.1 https://www.aigcaas.cn/product/detail/182.html
	ElldrethsDaydreamMixV10                   string = `elldrethsDaydreamMix_v10`                   // Elldreth's DayDream mix https://www.aigcaas.cn/product/detail/181.html
	TShirtPrintDesignsTestV01                 string = `tShirtPrintDesignsTest_v01`                 // T-shirt print designs (test model) https://www.aigcaas.cn/product/detail/180.html
	UrbandesignSd15V52                        string = `urbandesignSd15_v52`                        // urbandesign_sd15 https://www.aigcaas.cn/product/detail/179.html
	ElldrethsDreamMixV10                      string = `elldrethsDreamMix_v10`                      // Elldreth's Dream Mix https://www.aigcaas.cn/product/detail/178.html
	ElldrethsVividMixV20VividEr               string = `elldrethsVividMix_v20VividEr`               // Elldreth's Vivid Mix https://www.aigcaas.cn/product/detail/177.html
	HandpaintedRPGIconsV1                     string = `handpaintedRPGIcons_v1`                     // Handpainted RPG Icons https://www.aigcaas.cn/product/detail/176.html
	Dvmech200                                 string = `dvmech_200`                                 // dvMech https://www.aigcaas.cn/product/detail/175.html
	EpicrealismNewEra                         string = `epicrealism_newEra`                         // epiCRealism https://www.aigcaas.cn/product/detail/174.html
	MajicmixFantasyV20                        string = `majicmixFantasy_v20`                        // majicMIX fantasy https://www.aigcaas.cn/product/detail/173.html
	NightSkyYOZORAStyleYozoraV1Origin         string = `nightSkyYOZORAStyle_yozoraV1Origin`         // Night Sky YOZORA Style Model https://www.aigcaas.cn/product/detail/172.html
	ColorfulV30                               string = `colorful_v30`                               // Colorful https://www.aigcaas.cn/product/detail/171.html
	FantassifiedIconsFantassifiedIconsV20     string = `fantassifiedIcons_fantassifiedIconsV20`     // Fantassified Icons https://www.aigcaas.cn/product/detail/170.html
	SxzLuma099VAE                             string = `sxzLuma_099VAE`                             // SXZ Luma https://www.aigcaas.cn/product/detail/169.html
	BeenyouR13                                string = `beenyou_r13`                                // BeenYou https://www.aigcaas.cn/product/detail/168.html
	TmndMixTmndMixIVPruned                    string = `tmndMix_tmndMixIVPruned`                    // TMND-Mix https://www.aigcaas.cn/product/detail/167.html
	DisneyPixarCartoonV10                     string = `disneyPixarCartoon_v10`                     // Disney Pixar Cartoon Type A https://www.aigcaas.cn/product/detail/166.html
	LofiV22                                   string = `lofi_V22`                                   // LOFI https://www.aigcaas.cn/product/detail/165.html
	MajicmixLuxV10                            string = `majicmixLux_v10`                            // majicMIX lux https://www.aigcaas.cn/product/detail/164.html
	DivineelegancemixV5                       string = `divineelegancemix_V5`                       // DivineEleganceMix https://www.aigcaas.cn/product/detail/163.html
	Xxmix9realisticV30                        string = `xxmix9realistic_v30`                        // XXMix_9realistic https://www.aigcaas.cn/product/detail/162.html
	ComicBabesV1                              string = `comicBabes_v1`                              // Comic Babes https://www.aigcaas.cn/product/detail/161.html
	MaturemalemixV14                          string = `maturemalemix_v14`                          // maturemalemix https://www.aigcaas.cn/product/detail/160.html
	GameIconInstituteV30                      string = `gameIconInstitute_v30`                      // Game Icon Institute_mode https://www.aigcaas.cn/product/detail/159.html
	DarkSushi25D25DV20                        string = `darkSushi25D25D_v20`                        // Dark Sushi 2.5D 大颗寿司2.5D https://www.aigcaas.cn/product/detail/158.html
	OrientalMixV2V25                          string = `orientalMixV2_v25`                          // oriental_mix v2 https://www.aigcaas.cn/product/detail/157.html
	PixhellV20                                string = `pixhell_v20`                                // PIXHELL https://www.aigcaas.cn/product/detail/156.html
	BeautypromixV1                            string = `beautypromix_v1`                            // BeautyProMix https://www.aigcaas.cn/product/detail/155.html
	JuggernautV19                             string = `juggernaut_v19`                             // Juggernaut https://www.aigcaas.cn/product/detail/154.html
	CheeseDaddys41                            string = `cheeseDaddys_41`                            // Cheese Daddy's Landscapes mix https://www.aigcaas.cn/product/detail/153.html
	Abyssorangemix2SFWAbyssorangemix2Sfw      string = `abyssorangemix2SFW_abyssorangemix2Sfw`      // AbyssOrangeMix2 - SFW/Soft NSFW https://www.aigcaas.cn/product/detail/152.html
	ProtogenX53PhotorealismProtogenX53        string = `protogenX53Photorealism_protogenX53`        // Protogen x5.3 (Photorealism) Official Release https://www.aigcaas.cn/product/detail/151.html
	D526mixV145V145                           string = `526mixV145_v145`                            // 526Mix V1.4.5 https://www.aigcaas.cn/product/detail/150.html
	EtherRealMix21                            string = `etherRealMix_21`                            // Ether Real Mix https://www.aigcaas.cn/product/detail/149.html
	StoreBoughtGyozaMixV32                    string = `storeBoughtGyozaMix_v32`                    // Store Bought Gyoza (餃子Mix) https://www.aigcaas.cn/product/detail/148.html
	AnimatrixV20                              string = `animatrix_v20`                              // animatrix https://www.aigcaas.cn/product/detail/147.html
	ProfantasyV22                             string = `profantasy_v22`                             // ProFantasy https://www.aigcaas.cn/product/detail/146.html
	IlluminuttyDiffusionV111                  string = `illuminuttyDiffusion_v111`                  // Illuminutty Diffusion https://www.aigcaas.cn/product/detail/145.html
	VinteprotogenmixV20                       string = `vinteprotogenmix_V20`                       // VinteProtogenMix https://www.aigcaas.cn/product/detail/144.html
	PikasNewGenerationV20                     string = `pikasNewGeneration_v20`                     // Pika's New Generation https://www.aigcaas.cn/product/detail/143.html
	ProtogenInfinityProtogenInfinity          string = `protogenInfinity_protogenInfinity`          // Protogen Infinity Official Release https://www.aigcaas.cn/product/detail/142.html
	NoosphereV2                               string = `noosphere_v2`                               // Noosphere https://www.aigcaas.cn/product/detail/141.html
	QgoPromptingrealQgoPromptingrealV1        string = `qgoPromptingreal_qgoPromptingrealV1`        // QGO - PromptingReal https://www.aigcaas.cn/product/detail/140.html
	ToonyouBeta3                              string = `toonyou_beta3`                              // ToonYou https://www.aigcaas.cn/product/detail/139.html
	FantasyWorldV1                            string = `fantasyWorld_v1`                            // Fantasy World https://www.aigcaas.cn/product/detail/138.html
	HellmixV10                                string = `hellmix_v10`                                // HELLmix https://www.aigcaas.cn/product/detail/137.html
	AZovyaRPGArtistToolsV3                    string = `aZovyaRPGArtistTools_v3`                    // A-Zovya RPG Artist Tools https://www.aigcaas.cn/product/detail/136.html
	CheeseDaddysLandscapes35FP16              string = `cheeseDaddysLandscapes_35FP16`              // Cheese Daddy's Landscapes mix | OFFSET NOISE https://www.aigcaas.cn/product/detail/135.html
	RevMixV13                                 string = `revMix_v13`                                 // REV Mix https://www.aigcaas.cn/product/detail/134.html
	Deepboys25DV30                            string = `deepboys25D_v30`                            // DeepBoys_2.5D https://www.aigcaas.cn/product/detail/133.html
	YdenV3TrainfixYdenT                       string = `ydenV3Trainfix_ydenT`                       // yden V3 TrainFix https://www.aigcaas.cn/product/detail/132.html
	UltraILLUSI0N2DUltraILLUSI0N2DV1          string = `ultraILLUSI0N2D_ultraILLUSI0N2DV1`          // ULTRA-ILLUSI0N 2D https://www.aigcaas.cn/product/detail/131.html
	BeautifulRealisticBrav5                   string = `beautifulRealistic_brav5`                   // Beautiful Realistic Asians https://www.aigcaas.cn/product/detail/130.html
	AnythingAndAnythingAndEverything          string = `anythingAnd_anythingAndEverything`          // Anything and Everything https://www.aigcaas.cn/product/detail/129.html
	DualPersonalityDualEee                    string = `dualPersonality_dualEee`                    // dual_personality https://www.aigcaas.cn/product/detail/128.html
	HasdxHasedsdx                             string = `hasdx_hasedsdx`                             // HASDX https://www.aigcaas.cn/product/detail/127.html
	ViewerMixV17ViewerMixV17V2                string = `viewerMixV17_viewerMixV17V2`                // viewer-mix_v1.7 https://www.aigcaas.cn/product/detail/126.html
	AlstroemeriaMixV10                        string = `alstroemeriaMix_v10`                        // Alstroemeria Mix https://www.aigcaas.cn/product/detail/125.html
	Pastelboys2DV30                           string = `pastelboys2D_v30`                           // PastelBoys_2D https://www.aigcaas.cn/product/detail/124.html
	HolokukiV2                                string = `holokuki_v2`                                // HoloKuki https://www.aigcaas.cn/product/detail/123.html
	Samaritan3dCartoonV10                     string = `samaritan3dCartoon_v10`                     // Samaritan 3d Cartoon https://www.aigcaas.cn/product/detail/122.html
	PrmjV1                                    string = `prmj_v1`                                    // PRMJ https://www.aigcaas.cn/product/detail/121.html
	CetusMixV4                                string = `cetusMix_v4`                                // Cetus-Mix https://www.aigcaas.cn/product/detail/119.html
	D3Guofeng3V33                             string = `3Guofeng3_v33`                              // 国风3 GuoFeng3 https://www.aigcaas.cn/product/detail/118.html
	DarkSushiMixMixColorful                   string = `darkSushiMixMix_colorful`                   // Dark Sushi Mix 大颗寿司Mix https://www.aigcaas.cn/product/detail/117.html
	BreakdomainM2150                          string = `breakdomain_M2150`                          // BreakDomain https://www.aigcaas.cn/product/detail/116.html
	RealisianV40                              string = `realisian_v40`                              // Realisian https://www.aigcaas.cn/product/detail/115.html
	PulpV10                                   string = `pulp_v10`                                   // Pulp https://www.aigcaas.cn/product/detail/114.html
	AmixxV20Prunedfp16                        string = `amixx_v20Prunedfp16`                        // Amixx https://www.aigcaas.cn/product/detail/113.html
	EdgNendoV10                               string = `edgNendo_v10`                               // EDG&Nendo https://www.aigcaas.cn/product/detail/112.html
	LocsChinaLandscapesLocsChinaLandscapes    string = `locsChinaLandscapes_locsChinaLandscapes`    // Locs China Landscapes v2 https://www.aigcaas.cn/product/detail/111.html
	AniflatmixAnimeFlatColorStyleV20          string = `aniflatmixAnimeFlatColorStyle_v20`          // Aniflatmix - Anime Flat Color Style Mix (平涂り風/平涂风) https://www.aigcaas.cn/product/detail/110.html
	DndMapGeneratorV3                         string = `dndMapGenerator_v3`                         // DnD Map Generator https://www.aigcaas.cn/product/detail/109.html
	LattezenLandscapeV10                      string = `lattezenLandscape_v10`                      // lattezen_landscape https://www.aigcaas.cn/product/detail/108.html
	LandscapePhotorealV1                      string = `landscapePhotoreal_v1`                      // Landscape_PhotoReal https://www.aigcaas.cn/product/detail/107.html
	TrexmixV1                                 string = `trexmix_V1`                                 // Photography And Landscapes https://www.aigcaas.cn/product/detail/106.html
	CheesedOutAnimeV12                        string = `cheesedOutAnime_v12`                        // Cheesed out Anime Backgrounds https://www.aigcaas.cn/product/detail/105.html
	XsarchitecturalV11                        string = `xsarchitectural_v11`                        // XSarchitectural-InteriorDesign-ForXSLora https://www.aigcaas.cn/product/detail/104.html
	RaRenderV20                               string = `raRender_v20`                               // Ra-render/建筑渲染效果 https://www.aigcaas.cn/product/detail/103.html
	Xsarchitecturalv1comV12                   string = `xsarchitecturalv1com_v12`                   // XSarchitecturalV1Commercialbuildingrendering https://www.aigcaas.cn/product/detail/102.html
	Xsarchitecturalv2V20                      string = `xsarchitecturalv2_v20`                      // XSarchitecturalV2_AnyModernBuilding https://www.aigcaas.cn/product/detail/101.html
	AargArchitectureResV10                    string = `aargArchitectureRes_v10`                    // AARG-Architecture-Res https://www.aigcaas.cn/product/detail/100.html
	XswbsinglewoodenbuilV10                   string = `xswbsinglewoodenbuil_v10`                   // XSWBSingleWoodenBuildingV1 https://www.aigcaas.cn/product/detail/99.html
	IsometricFutureIsometricFutures10         string = `isometricFuture_isometricFutures10`         // Isometric Future https://www.aigcaas.cn/product/detail/98.html
	YqModernexclusiveV10                      string = `yqModernexclusive_v10`                      // YQ_ModernExclusive/现代风格专精 https://www.aigcaas.cn/product/detail/97.html
	RevAnimatedV122                           string = `revAnimated_v122`                           // ReV Animated https://www.aigcaas.cn/product/detail/96.html
	MirageGOLDENZABDKLBMirageZAB              string = `mirageGOLDENZABDKLB_mirageZAB`              // MIRAGE.GOLDEN.ZAB.DKLB https://www.aigcaas.cn/product/detail/95.html
	SimplyBeautifulV10                        string = `simplyBeautiful_v10`                        // Simply Beautiful https://www.aigcaas.cn/product/detail/94.html
	FkingScifiV3                              string = `fkingScifi_v3`                              // fking_scifi https://www.aigcaas.cn/product/detail/93.html
	Alladin3Alladin3                          string = `alladin3_alladin3`                          // Alladin.3 https://www.aigcaas.cn/product/detail/92.html
	ConceptSheetConceptSheetAlpha             string = `conceptSheet_conceptSheetAlpha`             // concept-sheet https://www.aigcaas.cn/product/detail/91.html
	Realdosmix                                string = `realdosmix_`                                // RealDosMix https://www.aigcaas.cn/product/detail/90.html
	CheckpointYesmixV16Original               string = `CheckpointYesmix_v16Original`               // 【Checkpoint】YesMix https://www.aigcaas.cn/product/detail/89.html
	DeliberateV2                              string = `deliberate_v2`                              // Deliberate https://www.aigcaas.cn/product/detail/88.html
	LyrielV16                                 string = `lyriel_v16`                                 // Lyriel https://www.aigcaas.cn/product/detail/87.html
	DreamlikePhotoreal20DreamlikePhotoreal20  string = `dreamlikePhotoreal20_dreamlikePhotoreal20`  // Realistic Vision V2.0 https://www.aigcaas.cn/product/detail/86.html
	MeinamixMeinaV10                          string = `meinamix_meinaV10`                          // MeinaMix https://www.aigcaas.cn/product/detail/85.html
	MarkerhandsketchV1                        string = `markerhandsketch-V1`                        // architectural design sketches with markers https://www.aigcaas.cn/product/detail/84.html
	Oxigien2ProSD21HiresV3Epic                string = `oxigien2ProSD21Hires_v3Epic`                // BreakDomainRealistic https://www.aigcaas.cn/product/detail/83.html
	TangYuanMixV20                            string = `tangYuanMix_v20`                            // Tang Yuan (汤圆Mix) https://www.aigcaas.cn/product/detail/82.html
	DvarchMultiPromptDvarchExterior           string = `dvarchMultiPrompt_dvarchExterior`           // dvArch - Multi-Prompt Architecture Tuned Model https://www.aigcaas.cn/product/detail/81.html
	D28DSTABLEBESTVERSIONRealTang             string = `28DSTABLEBESTVERSION_realTang`              // 2.8D STABLE BEST VERSION https://www.aigcaas.cn/product/detail/80.html
	Dreamshaper6BakedVae                      string = `dreamshaper_6BakedVae`                      // DreamShaper https://www.aigcaas.cn/product/detail/79.html
	ProtogenV22AnimeProtogenV22               string = `protogenV22Anime_protogenV22`               // Protogen v2.2 (Anime) Official Release https://www.aigcaas.cn/product/detail/78.html
	ProtogenX34PhotorealismProtogenX34        string = `protogenX34Photorealism_protogenX34`        // Protogen x3.4 (Photorealism) Official Release https://www.aigcaas.cn/product/detail/77.html
	CounterfeitV30V30                         string = `CounterfeitV30_v30`                         // Counterfeit-V3.0 https://www.aigcaas.cn/product/detail/76.html
	DdiconV10                                 string = `ddicon_v10`                                 // DDicon https://www.aigcaas.cn/product/detail/75.html
	ProductDesignEddiemauro15b                string = `productDesign_eddiemauro15b`                // Product Design (minimalism-eddiemauro) https://www.aigcaas.cn/product/detail/74.html
	Altdiffusion                              string = `altdiffusion`                               // BAAI/AltDiffusion https://www.aigcaas.cn/product/detail/70.html
)

// 人像卡通化（DCT-Net 模型）https://www.aigcaas.cn/product/detail/39.html
const (
	Dct                      string = `dct`
	DctMhfApiName            string = `mhf`             // 漫画风格
	DctWzryApiName           string = `wzry`            // 王者荣耀英雄风格
	DctHanddrawnApiName      string = `handdrawn`       // 手绘风格
	DctCompoundApiName       string = `compound`        // 通用风格
	Dct3DApiName             string = `3d`              // 3d风格
	DctArtstyleApiName       string = `artstyle`        // 艺术风格
	DctSketchApiName         string = `sketch`          // 素描风格
	DctDesignCompoundApiName string = `design_compound` // 插画风格
)

// 图像人脸融合 https://www.aigcaas.cn/product/detail/69.html
const (
	Facefu        string = `facefu`
	FacefuApiName string = `default`
)

// 语音识别（Paraformer） https://www.aigcaas.cn/product/detail/68.html
const (
	Paraformer        string = `paraformer`
	ParaformerApiName string = `default`
)

// 连续语义增强机器翻译（CSANMT） https://www.aigcaas.cn/product/detail/67.html
const (
	Csanmt        string = `csanmt`
	CsanmtApiName string = `default`
)

// 图像去噪与去模糊化（NAFNet） https://www.aigcaas.cn/product/detail/63.html
const (
	MafnetDenoise        string = `mafnet_denoise`
	MafnetDenoiseApiName string = `denoise`
)

// 图像去摩尔纹（uhdm） https://www.aigcaas.cn/product/detail/61.html
const (
	Uhdm        string = `uhdm`
	UhdmApiName string = `default`
)

// 图像去色带（RealESRGAN） https://www.aigcaas.cn/product/detail/60.html
const (
	Rrdb        string = `rrdb`
	RrdbApiName string = `default`
)

// 垃圾分类（ConvNeXt） https://www.aigcaas.cn/product/detail/58.html
const (
	Convnext        string = `convnext`
	ConvnextApiName string = `default`
)

// 实时检测（Yolo） https://www.aigcaas.cn/product/detail/57.html
const (
	RealtimeDetection            string = `realtime_detection`
	RealtimeDetectionTrafficSign string = `traffic_sign` // 交通标识检测
	RealtimeDetectionCigarette   string = `cigarette`    // 香烟检测
	RealtimeDetectionSafetyHat   string = `safety_hat`   // 安全帽检测
	RealtimeDetectionFacemask    string = `facemask`     // 口罩检测
	RealtimeDetectionHand        string = `hand`         // 手部检测
	RealtimeDetectionHead        string = `head`         // 人头检测
	RealtimeDetectionHuman       string = `human`        // 人体检测
	RealtimeDetectionPhone       string = `phone`        // 手机检测
)

// 二次元风格生成扩散模型（anything） https://www.aigcaas.cn/product/detail/56.html
const (
	Anything                   string = `anything`
	AnythingImg2imgV45ApiName  string = `img2imgv45`  // 图生图（4.5）
	AnythingImg2imgV40ApiName  string = `img2imgv40`  // 图生图（4.0）
	AnythingText2imgV45ApiName string = `text2imgv45` // 图生图（4.0）
	AnythingText2imgV40ApiName string = `text2imgv40` // 图生图（4.0）
)

// 行人结构化属性识别（ResNet50） https://www.aigcaas.cn/product/detail/55.html
const (
	Resnet50Pedestrian        string = `resnet50_pedestrian`
	Resnet50PedestrianApiName string = `default`
)

// ChatGLM 中英对话大模型 https://www.aigcaas.cn/product/detail/53.html
const (
	Chatglm        string = `chatglm`
	ChatglmApiName string = `default`
)

// 万物识别（Resnest101） https://www.aigcaas.cn/product/detail/46.html
const (
	GeneralRecognition        string = `general_recognition`
	GeneralRecognitionApiName string = `general`
)

// 图像上色（DeOldify） https://www.aigcaas.cn/product/detail/43.html
const (
	Deoldify        string = `deoldify`
	DeoldifyApiName string = `image`
)

// 人像修复增强（GPEN） https://www.aigcaas.cn/product/detail/42.html
const (
	Gpen        string = `gpen`
	GpenApiName string = `image`
)

// 图像上色模型（DDColor） https://www.aigcaas.cn/product/detail/41.html
const (
	Ddcolor        string = `ddcolor`
	DdcolorApiName string = `image`
)

// 图像调色（CSRNet） https://www.aigcaas.cn/product/detail/40.html
const (
	Csrnet        string = `csrnet`
	CsrnetApiName string = `color`
)

// 图像目标检测（ImageAI） https://www.aigcaas.cn/product/detail/37.html
const (
	Imageai        string = `imageai`
	ImageaiApiName string = `detection`
)

// 图像描述（vit-gpt2） https://www.aigcaas.cn/product/detail/33.html
const (
	Vitgpt2        string = `vitgpt2`
	Vitgpt2ApiName string = `desc`
)

// OCR 文字识别 https://www.aigcaas.cn/product/detail/29.html
const (
	Ocr              string = `ocr`
	OcrApiNameSource string = `source` // 原始模拟
	OcrApiNameSimple string = `simple` // 简单模式
)
