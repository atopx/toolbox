package startup

import "toolbox/common/system"

func init() {
	// api服务
	system.SetService(&WebapiService{})

	// TODO 其他任务，如异步任务
}
