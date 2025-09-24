package okxdexapi

type Code string

const SuccessCode Code = "0"

// Authentication and Authorization Error Codes (50xxx)
const (
	CodeServiceUnavailable  Code = "50001" // Service temporarily unavailable. Try again later
	CodeRateLimit           Code = "50011" // 用户请求频率过快，超过该接口允许的限额 / Rate limit reached
	CodeMissingParameter    Code = "50014" // 必填参数不能为空 / Parameter cannot be empty
	CodeSystemError         Code = "50026" // 系统错误，请稍后重试 / System error. Try again later
	CodeMissingAccessKey    Code = "50103" // 请求头"OK_ACCESS_KEY"不能为空 / Request header "OK-ACCESS-KEY" cannot be empty
	CodeMissingPassphrase   Code = "50104" // 请求头"OK_ACCESS_PASSPHRASE"不能为空 / Request header "OK-ACCESS-PASSPHRASE" cannot be empty
	CodeIncorrectPassphrase Code = "50105" // 请求头"OK_ACCESS_PASSPHRASE"错误 / Request header "OK-ACCESS-PASSPHRASE" incorrect
	CodeMissingSign         Code = "50106" // 请求头"OK_ACCESS_SIGN"不能为空 / Request header "OK-ACCESS-SIGN" cannot be empty
	CodeMissingTimestamp    Code = "50107" // 请求头"OK_ACCESS_TIMESTAMP"不能为空 / Request header "OK-ACCESS-TIMESTAMP" cannot be empty
	CodeInvalidAccessKey    Code = "50111" // 无效的 OK_ACCESS_KEY / Invalid OK-ACCESS-KEY
	CodeInvalidTimestamp    Code = "50112" // 无效的 OK_ACCESS_TIMESTAMP / Invalid OK-ACCESS-TIMESTAMP
	CodeInvalidSignature    Code = "50113" // 无效的签名 / Invalid signature
	CodeParameterError      Code = "51000" // 参数错误 / Parameter error
)

// Business Logic Error Codes (80xxx)
const (
	CodeDuplicateSubmission      Code = "80000" // 重复提交
	CodeCallDataExceedsLimit     Code = "80001" // callData 超出最大限制，请在 5 分钟后重试
	CodeTokenObjectLimitReached  Code = "80002" // 请求的代币 Object 数量达到上限
	CodeMainnetTokenLimitReached Code = "80003" // 请求的主网代币 Object 数量达到上限
	CodeSUIObjectQueryTimeout    Code = "80004" // 查询 SUI Object 超时
	CodeInsufficientSUIObject    Code = "80005" // sui object 不足
)

// System and Chain Error Codes (81xxx)
const (
	CodeIncorrectParameter Code = "81001" // 错误参数 / Incorrect parameter
	CodeChainNotSupport    Code = "81104" // 链不支持 / Chain not support
	CodeWalletTypeMismatch Code = "81108" // Wallet type does not match the required type
	CodeCoinNotExist       Code = "81152" // Coin not exist
	CodeNodeReturnFailed   Code = "81451" // Node return failed
)

// Trading and Liquidity Error Codes (82xxx)
const (
	CodeInsufficientLiquidity    Code = "82000" // 流动性不足 / 此报价的流动性过低
	CodeCommissionServiceDown    Code = "82001" // 分佣服务暂不可用 / 报价未超过挂单的最小规模
	CodeOrderCannotProcess       Code = "82002" // 挂单无法处理此报价
	CodeInvalidCommissionAddress Code = "82003" // toToken 分佣地址有误 / 挂单拒绝响应此用户
	CodeFourMemeNotSupported     Code = "82004" // 通过 Four.meme 协议兑换的交易不支持分佣
	CodeAspectaNotSupported      Code = "82005" // 通过 Aspecta 协议兑换的交易不支持分佣
	CodeBelowMinCrossChain       Code = "82102" // 小于最小跨链数,最小数量是 {0}
	CodeAboveMaxCrossChain       Code = "82103" // 超出最大跨链数,最大数量是 {0}
	CodeUnsupportedCurrency      Code = "82104" // 不支持该币种
	CodeUnsupportedChain         Code = "82105" // 不支持该链
	CodeExcessivePriceImpact     Code = "82112" // 当前交易的询价路径价差超过限制，可能造成用户的资产损失
	CodeCallDataLimitRetry       Code = "82116" // callData 超出最大限制，请在 5 分钟后重试
	CodeNoAuthorizationNeeded    Code = "82130" // 链不需要进行授权交易，可直接兑换
)

// ErrorInfo 错误信息结构体
type ErrorInfo struct {
	Code       Code   // 错误码
	Message    string // 中文错误消息
	MessageEN  string // 英文错误消息
	Category   string // 错误分类
	HTTPStatus int    // 建议的HTTP状态码
	Retryable  bool   // 是否可重试
}

var errorMap = map[Code]ErrorInfo{
	// 认证授权错误 (50xxx)
	CodeServiceUnavailable: {
		Code:       CodeServiceUnavailable,
		Message:    "服务暂时不可用，请稍后重试",
		MessageEN:  "Service temporarily unavailable. Try again later",
		Category:   "ServiceError",
		HTTPStatus: 503,
		Retryable:  true,
	},
	CodeRateLimit: {
		Code:       CodeRateLimit,
		Message:    "用户请求频率过快，超过该接口允许的限额",
		MessageEN:  "Rate limit reached",
		Category:   "RateLimit",
		HTTPStatus: 429,
		Retryable:  true,
	},
	CodeMissingParameter: {
		Code:       CodeMissingParameter,
		Message:    "必填参数不能为空",
		MessageEN:  "Parameter cannot be empty",
		Category:   "ValidationError",
		HTTPStatus: 400,
		Retryable:  false,
	},
	CodeSystemError: {
		Code:       CodeSystemError,
		Message:    "系统错误，请稍后重试",
		MessageEN:  "System error. Try again later",
		Category:   "SystemError",
		HTTPStatus: 500,
		Retryable:  true,
	},
	CodeMissingAccessKey: {
		Code:       CodeMissingAccessKey,
		Message:    "请求头OK_ACCESS_KEY不能为空",
		MessageEN:  "Request header OK-ACCESS-KEY cannot be empty",
		Category:   "AuthError",
		HTTPStatus: 401,
		Retryable:  false,
	},
	CodeMissingPassphrase: {
		Code:       CodeMissingPassphrase,
		Message:    "请求头OK_ACCESS_PASSPHRASE不能为空",
		MessageEN:  "Request header OK-ACCESS-PASSPHRASE cannot be empty",
		Category:   "AuthError",
		HTTPStatus: 401,
		Retryable:  false,
	},
	CodeIncorrectPassphrase: {
		Code:       CodeIncorrectPassphrase,
		Message:    "请求头OK_ACCESS_PASSPHRASE错误",
		MessageEN:  "Request header OK-ACCESS-PASSPHRASE incorrect",
		Category:   "AuthError",
		HTTPStatus: 401,
		Retryable:  false,
	},
	CodeMissingSign: {
		Code:       CodeMissingSign,
		Message:    "请求头OK_ACCESS_SIGN不能为空",
		MessageEN:  "Request header OK-ACCESS-SIGN cannot be empty",
		Category:   "AuthError",
		HTTPStatus: 401,
		Retryable:  false,
	},
	CodeMissingTimestamp: {
		Code:       CodeMissingTimestamp,
		Message:    "请求头OK_ACCESS_TIMESTAMP不能为空",
		MessageEN:  "Request header OK-ACCESS-TIMESTAMP cannot be empty",
		Category:   "AuthError",
		HTTPStatus: 401,
		Retryable:  false,
	},
	CodeInvalidAccessKey: {
		Code:       CodeInvalidAccessKey,
		Message:    "无效的OK_ACCESS_KEY",
		MessageEN:  "Invalid OK-ACCESS-KEY",
		Category:   "AuthError",
		HTTPStatus: 401,
		Retryable:  false,
	},
	CodeInvalidTimestamp: {
		Code:       CodeInvalidTimestamp,
		Message:    "无效的OK_ACCESS_TIMESTAMP",
		MessageEN:  "Invalid OK-ACCESS-TIMESTAMP",
		Category:   "AuthError",
		HTTPStatus: 401,
		Retryable:  false,
	},
	CodeInvalidSignature: {
		Code:       CodeInvalidSignature,
		Message:    "无效的签名",
		MessageEN:  "Invalid signature",
		Category:   "AuthError",
		HTTPStatus: 401,
		Retryable:  false,
	},
	CodeParameterError: {
		Code:       CodeParameterError,
		Message:    "参数错误",
		MessageEN:  "Parameter error",
		Category:   "ValidationError",
		HTTPStatus: 400,
		Retryable:  false,
	},

	// 业务逻辑错误 (80xxx)
	CodeDuplicateSubmission: {
		Code:       CodeDuplicateSubmission,
		Message:    "重复提交",
		MessageEN:  "Duplicate submission",
		Category:   "BusinessError",
		HTTPStatus: 409,
		Retryable:  false,
	},
	CodeCallDataExceedsLimit: {
		Code:       CodeCallDataExceedsLimit,
		Message:    "callData 超出最大限制，请在 5 分钟后重试",
		MessageEN:  "callData exceeds maximum limit, please retry after 5 minutes",
		Category:   "LimitError",
		HTTPStatus: 429,
		Retryable:  true,
	},
	CodeTokenObjectLimitReached: {
		Code:       CodeTokenObjectLimitReached,
		Message:    "请求的代币 Object 数量达到上限",
		MessageEN:  "Token object request limit reached",
		Category:   "LimitError",
		HTTPStatus: 429,
		Retryable:  true,
	},
	CodeMainnetTokenLimitReached: {
		Code:       CodeMainnetTokenLimitReached,
		Message:    "请求的主网代币 Object 数量达到上限",
		MessageEN:  "Mainnet token object request limit reached",
		Category:   "LimitError",
		HTTPStatus: 429,
		Retryable:  true,
	},
	CodeSUIObjectQueryTimeout: {
		Code:       CodeSUIObjectQueryTimeout,
		Message:    "查询 SUI Object 超时",
		MessageEN:  "SUI Object query timeout",
		Category:   "TimeoutError",
		HTTPStatus: 408,
		Retryable:  true,
	},
	CodeInsufficientSUIObject: {
		Code:       CodeInsufficientSUIObject,
		Message:    "sui object 不足",
		MessageEN:  "Insufficient SUI object",
		Category:   "InsufficientError",
		HTTPStatus: 400,
		Retryable:  false,
	},

	// 系统和链错误 (81xxx)
	CodeIncorrectParameter: {
		Code:       CodeIncorrectParameter,
		Message:    "错误参数",
		MessageEN:  "Incorrect parameter",
		Category:   "ValidationError",
		HTTPStatus: 400,
		Retryable:  false,
	},
	CodeChainNotSupport: {
		Code:       CodeChainNotSupport,
		Message:    "链不支持",
		MessageEN:  "Chain not support",
		Category:   "UnsupportedError",
		HTTPStatus: 400,
		Retryable:  false,
	},
	CodeWalletTypeMismatch: {
		Code:       CodeWalletTypeMismatch,
		Message:    "钱包类型与所需类型不匹配",
		MessageEN:  "Wallet type does not match the required type",
		Category:   "ValidationError",
		HTTPStatus: 400,
		Retryable:  false,
	},
	CodeCoinNotExist: {
		Code:       CodeCoinNotExist,
		Message:    "代币不存在",
		MessageEN:  "Coin not exist",
		Category:   "NotFoundError",
		HTTPStatus: 404,
		Retryable:  false,
	},
	CodeNodeReturnFailed: {
		Code:       CodeNodeReturnFailed,
		Message:    "节点返回失败",
		MessageEN:  "Node return failed",
		Category:   "NetworkError",
		HTTPStatus: 502,
		Retryable:  true,
	},

	// 交易和流动性错误 (82xxx)
	CodeInsufficientLiquidity: {
		Code:       CodeInsufficientLiquidity,
		Message:    "流动性不足，此报价的流动性过低",
		MessageEN:  "Insufficient liquidity",
		Category:   "LiquidityError",
		HTTPStatus: 400,
		Retryable:  false,
	},
	CodeCommissionServiceDown: {
		Code:       CodeCommissionServiceDown,
		Message:    "分佣服务暂不可用，报价未超过挂单的最小规模",
		MessageEN:  "Commission service temporarily unavailable",
		Category:   "ServiceError",
		HTTPStatus: 503,
		Retryable:  true,
	},
	CodeOrderCannotProcess: {
		Code:       CodeOrderCannotProcess,
		Message:    "挂单无法处理此报价",
		MessageEN:  "Order cannot process this quote",
		Category:   "OrderError",
		HTTPStatus: 400,
		Retryable:  false,
	},
	CodeInvalidCommissionAddress: {
		Code:       CodeInvalidCommissionAddress,
		Message:    "toToken 分佣地址有误，挂单拒绝响应此用户",
		MessageEN:  "Invalid commission address",
		Category:   "ValidationError",
		HTTPStatus: 400,
		Retryable:  false,
	},
	CodeFourMemeNotSupported: {
		Code:       CodeFourMemeNotSupported,
		Message:    "通过 Four.meme 协议兑换的交易不支持分佣",
		MessageEN:  "Four.meme protocol transactions do not support commission",
		Category:   "UnsupportedError",
		HTTPStatus: 400,
		Retryable:  false,
	},
	CodeAspectaNotSupported: {
		Code:       CodeAspectaNotSupported,
		Message:    "通过 Aspecta 协议兑换的交易不支持分佣",
		MessageEN:  "Aspecta protocol transactions do not support commission",
		Category:   "UnsupportedError",
		HTTPStatus: 400,
		Retryable:  false,
	},
	CodeBelowMinCrossChain: {
		Code:       CodeBelowMinCrossChain,
		Message:    "小于最小跨链数",
		MessageEN:  "Below minimum cross-chain amount",
		Category:   "ValidationError",
		HTTPStatus: 400,
		Retryable:  false,
	},
	CodeAboveMaxCrossChain: {
		Code:       CodeAboveMaxCrossChain,
		Message:    "超出最大跨链数",
		MessageEN:  "Above maximum cross-chain amount",
		Category:   "ValidationError",
		HTTPStatus: 400,
		Retryable:  false,
	},
	CodeUnsupportedCurrency: {
		Code:       CodeUnsupportedCurrency,
		Message:    "不支持该币种",
		MessageEN:  "Unsupported currency",
		Category:   "UnsupportedError",
		HTTPStatus: 400,
		Retryable:  false,
	},
	CodeUnsupportedChain: {
		Code:       CodeUnsupportedChain,
		Message:    "不支持该链",
		MessageEN:  "Unsupported chain",
		Category:   "UnsupportedError",
		HTTPStatus: 400,
		Retryable:  false,
	},
	CodeExcessivePriceImpact: {
		Code:       CodeExcessivePriceImpact,
		Message:    "当前交易的询价路径价差超过限制，可能造成用户的资产损失",
		MessageEN:  "Excessive price impact detected",
		Category:   "PriceError",
		HTTPStatus: 400,
		Retryable:  false,
	},
	CodeCallDataLimitRetry: {
		Code:       CodeCallDataLimitRetry,
		Message:    "callData 超出最大限制，请在 5 分钟后重试",
		MessageEN:  "callData exceeds maximum limit, please retry after 5 minutes",
		Category:   "LimitError",
		HTTPStatus: 429,
		Retryable:  true,
	},
	CodeNoAuthorizationNeeded: {
		Code:       CodeNoAuthorizationNeeded,
		Message:    "链不需要进行授权交易，可直接兑换",
		MessageEN:  "Chain does not require authorization transaction, can exchange directly",
		Category:   "InfoMessage",
		HTTPStatus: 200,
		Retryable:  false,
	},
}

// GetErrorInfo 根据错误码获取完整的错误信息
func GetErrorInfo(code Code) ErrorInfo {

	// 如果找到对应的错误信息，返回它；否则返回未知错误
	if info, exists := errorMap[code]; exists {
		return info
	}

	// 未知错误码的默认处理
	return ErrorInfo{
		Code:       code,
		Message:    "未知错误码: " + string(code),
		MessageEN:  "Unknown error code: " + string(code),
		Category:   "UnknownError",
		HTTPStatus: 500,
		Retryable:  false,
	}
}

// GetErrorMessage 获取错误的中文消息
func GetErrorMessage(code Code) string {
	return GetErrorInfo(code).Message
}

// GetErrorMessageEN 获取错误的英文消息
func GetErrorMessageEN(code Code) string {
	return GetErrorInfo(code).MessageEN
}

// IsRetryable 判断错误是否可重试
func IsRetryable(code Code) bool {
	return GetErrorInfo(code).Retryable
}

// GetErrorCategory 获取错误分类
func GetErrorCategory(code Code) string {
	return GetErrorInfo(code).Category
}

// GetHTTPStatus 获取建议的HTTP状态码
func GetHTTPStatus(code Code) int {
	return GetErrorInfo(code).HTTPStatus
}

// IsSuccess 判断是否为成功码
func IsSuccess(code Code) bool {
	return code == SuccessCode
}

// IsAuthError 判断是否为认证错误
func IsAuthError(code Code) bool {
	return GetErrorCategory(code) == "AuthError"
}

// IsValidationError 判断是否为参数验证错误
func IsValidationError(code Code) bool {
	return GetErrorCategory(code) == "ValidationError"
}

// IsSystemError 判断是否为系统错误
func IsSystemError(code Code) bool {
	category := GetErrorCategory(code)
	return category == "SystemError" || category == "ServiceError" || category == "NetworkError"
}

// IsBusinessError 判断是否为业务逻辑错误
func IsBusinessError(code Code) bool {
	category := GetErrorCategory(code)
	return category == "BusinessError" || category == "LiquidityError" || category == "OrderError"
}
