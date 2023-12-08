package version

// 版本参数
var (
	gitTag     string
	gitCommit  string
	gitBranch  string
	buildTime  string
	goVersion  string
	appVersion string
	appName    string
)

// GetGitTag git仓库tag标签
// $(shell git log --pretty=format:"%ad_%h" -1 --date=short)
func GetGitTag() string {
	return gitTag
}

// GetGitCommit git提交commitID
// $(shell git rev-parse HEAD)
func GetGitCommit() string {
	return gitCommit
}

// GetGitBranch git提交版本号
// $(shell git symbolic-ref --short -q HEAD)
func GetGitBranch() string {
	return gitBranch
}

// GetBuildTime 编译时间
// $(shell date "+%F %T")
func GetBuildTime() string {
	return buildTime
}

// GetGoVersion go版本信息
// $(shell go version)
func GetGoVersion() string {
	return goVersion
}

// GetAppVersion 应用版本号
func GetAppVersion() string {
	return appVersion
}

// GetAppName 应用标志，后续作用于日志，trace，metric中
func GetAppName() string {
	return appName
}
