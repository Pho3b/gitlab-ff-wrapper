package main

import (
	"fmt"
	"github.com/pho3b/gitlab-ff-wrapper/constants"
	"github.com/pho3b/gitlab-ff-wrapper/enums"
	"github.com/pho3b/gitlab-ff-wrapper/ffclient"
	"github.com/pho3b/gitlab-ff-wrapper/fflags"
	"os"
)

func main() {
	os.Setenv(constants.EnvTypeVariableName, enums.Development.ToString())
	ffclient.Init(
		"https://gitlab.com/api/v4/feature_flags/unleash/67579289",
		"glffct-Lcr77RVN5RRj2rRv71Jt",
	)

	fmt.Println(ffclient.Get().IsFeatureEnabled(fflags.MyFooFeatureFlag))

	//os.Setenv(constants.EnvTypeVariableName, enums.Development.ToString())
	//ffclient.InitWithConfig(
	//	ffconfig.ClientConfig{
	//		ProjectUrl:                  "https://gitlab.com/api/v4/feature_flags/unleash/67579289",
	//		ProjectId:                   "glffct-Lcr77RVN5RRj2rRv71Jt",
	//		Logger:                      nil,
	//		ValidEnvironmentTypes:       nil,
	//		AsyncInitialization:         false,
	//		EnvironmentTypeVariableName: "",
	//	},
	//)
	//
	//fmt.Println(ffclient.Get().IsFeatureEnabled(fflags.MyFooFeatureFlag))
}
