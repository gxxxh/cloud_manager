package main

import cloud_manager "cloud_manager/src/analyzer"

func main() {
	analyzer, _ := cloud_manager.NewCloudAPIAnalyzer("aliyun")
	analyzer.ExtractCloudAPIs()
}
