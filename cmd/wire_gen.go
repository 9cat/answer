// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package answercmd

import (
	"github.com/answerdev/answer/internal/base/conf"
	"github.com/answerdev/answer/internal/base/cron"
	"github.com/answerdev/answer/internal/base/data"
	"github.com/answerdev/answer/internal/base/middleware"
	"github.com/answerdev/answer/internal/base/server"
	"github.com/answerdev/answer/internal/base/translator"
	"github.com/answerdev/answer/internal/controller"
	"github.com/answerdev/answer/internal/controller/template_render"
	"github.com/answerdev/answer/internal/controller_admin"
	"github.com/answerdev/answer/internal/repo/activity"
	"github.com/answerdev/answer/internal/repo/activity_common"
	"github.com/answerdev/answer/internal/repo/answer"
	"github.com/answerdev/answer/internal/repo/auth"
	"github.com/answerdev/answer/internal/repo/captcha"
	"github.com/answerdev/answer/internal/repo/collection"
	"github.com/answerdev/answer/internal/repo/comment"
	"github.com/answerdev/answer/internal/repo/config"
	"github.com/answerdev/answer/internal/repo/export"
	"github.com/answerdev/answer/internal/repo/meta"
	"github.com/answerdev/answer/internal/repo/notification"
	"github.com/answerdev/answer/internal/repo/plugin_config"
	"github.com/answerdev/answer/internal/repo/question"
	"github.com/answerdev/answer/internal/repo/rank"
	"github.com/answerdev/answer/internal/repo/reason"
	"github.com/answerdev/answer/internal/repo/report"
	"github.com/answerdev/answer/internal/repo/revision"
	"github.com/answerdev/answer/internal/repo/role"
	"github.com/answerdev/answer/internal/repo/search_common"
	"github.com/answerdev/answer/internal/repo/site_info"
	"github.com/answerdev/answer/internal/repo/tag"
	"github.com/answerdev/answer/internal/repo/tag_common"
	"github.com/answerdev/answer/internal/repo/unique"
	"github.com/answerdev/answer/internal/repo/user"
	"github.com/answerdev/answer/internal/repo/user_external_login"
	"github.com/answerdev/answer/internal/router"
	"github.com/answerdev/answer/internal/service"
	"github.com/answerdev/answer/internal/service/action"
	activity2 "github.com/answerdev/answer/internal/service/activity"
	activity_common2 "github.com/answerdev/answer/internal/service/activity_common"
	"github.com/answerdev/answer/internal/service/activity_queue"
	"github.com/answerdev/answer/internal/service/answer_common"
	auth2 "github.com/answerdev/answer/internal/service/auth"
	"github.com/answerdev/answer/internal/service/collection_common"
	comment2 "github.com/answerdev/answer/internal/service/comment"
	"github.com/answerdev/answer/internal/service/comment_common"
	config2 "github.com/answerdev/answer/internal/service/config"
	"github.com/answerdev/answer/internal/service/dashboard"
	export2 "github.com/answerdev/answer/internal/service/export"
	"github.com/answerdev/answer/internal/service/follow"
	meta2 "github.com/answerdev/answer/internal/service/meta"
	"github.com/answerdev/answer/internal/service/notice_queue"
	notification2 "github.com/answerdev/answer/internal/service/notification"
	"github.com/answerdev/answer/internal/service/notification_common"
	"github.com/answerdev/answer/internal/service/object_info"
	"github.com/answerdev/answer/internal/service/plugin_common"
	"github.com/answerdev/answer/internal/service/question_common"
	rank2 "github.com/answerdev/answer/internal/service/rank"
	reason2 "github.com/answerdev/answer/internal/service/reason"
	report2 "github.com/answerdev/answer/internal/service/report"
	"github.com/answerdev/answer/internal/service/report_admin"
	"github.com/answerdev/answer/internal/service/report_handle_admin"
	"github.com/answerdev/answer/internal/service/revision_common"
	role2 "github.com/answerdev/answer/internal/service/role"
	"github.com/answerdev/answer/internal/service/search_parser"
	"github.com/answerdev/answer/internal/service/service_config"
	"github.com/answerdev/answer/internal/service/siteinfo"
	"github.com/answerdev/answer/internal/service/siteinfo_common"
	tag2 "github.com/answerdev/answer/internal/service/tag"
	tag_common2 "github.com/answerdev/answer/internal/service/tag_common"
	"github.com/answerdev/answer/internal/service/uploader"
	"github.com/answerdev/answer/internal/service/user_admin"
	"github.com/answerdev/answer/internal/service/user_common"
	user_external_login2 "github.com/answerdev/answer/internal/service/user_external_login"
	"github.com/segmentfault/pacman"
	"github.com/segmentfault/pacman/log"
)

// Injectors from wire.go:

// initApplication init application.
func initApplication(debug bool, serverConf *conf.Server, dbConf *data.Database, cacheConf *data.CacheConf, i18nConf *translator.I18n, swaggerConf *router.SwaggerConfig, serviceConf *service_config.ServiceConfig, logConf log.Logger) (*pacman.Application, func(), error) {
	staticRouter := router.NewStaticRouter(serviceConf)
	i18nTranslator, err := translator.NewTranslator(i18nConf)
	if err != nil {
		return nil, nil, err
	}
	engine, err := data.NewDB(debug, dbConf)
	if err != nil {
		return nil, nil, err
	}
	cache, cleanup, err := data.NewCache(cacheConf)
	if err != nil {
		return nil, nil, err
	}
	dataData, cleanup2, err := data.NewData(engine, cache)
	if err != nil {
		cleanup()
		return nil, nil, err
	}
	siteInfoRepo := site_info.NewSiteInfo(dataData)
	siteInfoCommonService := siteinfo_common.NewSiteInfoCommonService(siteInfoRepo)
	langController := controller.NewLangController(i18nTranslator, siteInfoCommonService)
	authRepo := auth.NewAuthRepo(dataData)
	authService := auth2.NewAuthService(authRepo)
	userRepo := user.NewUserRepo(dataData)
	uniqueIDRepo := unique.NewUniqueIDRepo(dataData)
	configRepo := config.NewConfigRepo(dataData)
	configService := config2.NewConfigService(configRepo)
	activityRepo := activity_common.NewActivityRepo(dataData, uniqueIDRepo, configService)
	userRankRepo := rank.NewUserRankRepo(dataData, configService)
	userActiveActivityRepo := activity.NewUserActiveActivityRepo(dataData, activityRepo, userRankRepo, configService)
	emailRepo := export.NewEmailRepo(dataData)
	emailService := export2.NewEmailService(configService, emailRepo, siteInfoRepo)
	userRoleRelRepo := role.NewUserRoleRelRepo(dataData)
	roleRepo := role.NewRoleRepo(dataData)
	roleService := role2.NewRoleService(roleRepo)
	userRoleRelService := role2.NewUserRoleRelService(userRoleRelRepo, roleService)
	userCommon := usercommon.NewUserCommon(userRepo, userRoleRelService, authService, siteInfoCommonService)
	userExternalLoginRepo := user_external_login.NewUserExternalLoginRepo(dataData)
	userExternalLoginService := user_external_login2.NewUserExternalLoginService(userRepo, userCommon, userExternalLoginRepo, emailService, siteInfoCommonService, userActiveActivityRepo)
	userService := service.NewUserService(userRepo, userActiveActivityRepo, activityRepo, emailService, authService, siteInfoCommonService, userRoleRelService, userCommon, userExternalLoginService)
	captchaRepo := captcha.NewCaptchaRepo(dataData)
	captchaService := action.NewCaptchaService(captchaRepo)
	userController := controller.NewUserController(authService, userService, captchaService, emailService, siteInfoCommonService)
	commentRepo := comment.NewCommentRepo(dataData, uniqueIDRepo)
	commentCommonRepo := comment.NewCommentCommonRepo(dataData, uniqueIDRepo)
	answerRepo := answer.NewAnswerRepo(dataData, uniqueIDRepo, userRankRepo, activityRepo)
	questionRepo := question.NewQuestionRepo(dataData, uniqueIDRepo)
	tagCommonRepo := tag_common.NewTagCommonRepo(dataData, uniqueIDRepo)
	tagRelRepo := tag.NewTagRelRepo(dataData, uniqueIDRepo)
	tagRepo := tag.NewTagRepo(dataData, uniqueIDRepo)
	revisionRepo := revision.NewRevisionRepo(dataData, uniqueIDRepo)
	revisionService := revision_common.NewRevisionService(revisionRepo, userRepo)
	activityQueueService := activity_queue.NewActivityQueueService()
	tagCommonService := tag_common2.NewTagCommonService(tagCommonRepo, tagRelRepo, tagRepo, revisionService, siteInfoCommonService, activityQueueService)
	objService := object_info.NewObjService(answerRepo, questionRepo, commentCommonRepo, tagCommonRepo, tagCommonService)
	voteRepo := activity_common.NewVoteRepo(dataData, activityRepo)
	notificationQueueService := notice_queue.NewNotificationQueueService()
	commentService := comment2.NewCommentService(commentRepo, commentCommonRepo, userCommon, objService, voteRepo, emailService, userRepo, notificationQueueService, activityQueueService)
	rolePowerRelRepo := role.NewRolePowerRelRepo(dataData)
	rolePowerRelService := role2.NewRolePowerRelService(rolePowerRelRepo, userRoleRelService)
	rankService := rank2.NewRankService(userCommon, userRankRepo, objService, userRoleRelService, rolePowerRelService, configService)
	commentController := controller.NewCommentController(commentService, rankService, captchaService)
	reportRepo := report.NewReportRepo(dataData, uniqueIDRepo)
	reportService := report2.NewReportService(reportRepo, objService)
	reportController := controller.NewReportController(reportService, rankService, captchaService)
	serviceVoteRepo := activity.NewVoteRepo(dataData, activityRepo, userRankRepo, notificationQueueService)
	voteService := service.NewVoteService(serviceVoteRepo, configService, questionRepo, answerRepo, commentCommonRepo, objService)
	voteController := controller.NewVoteController(voteService, rankService, captchaService)
	followRepo := activity_common.NewFollowRepo(dataData, uniqueIDRepo, activityRepo)
	tagService := tag2.NewTagService(tagRepo, tagCommonService, revisionService, followRepo, siteInfoCommonService, activityQueueService)
	tagController := controller.NewTagController(tagService, tagCommonService, rankService)
	followFollowRepo := activity.NewFollowRepo(dataData, uniqueIDRepo, activityRepo)
	followService := follow.NewFollowService(followFollowRepo, followRepo, tagCommonRepo)
	followController := controller.NewFollowController(followService)
	collectionRepo := collection.NewCollectionRepo(dataData, uniqueIDRepo)
	collectionGroupRepo := collection.NewCollectionGroupRepo(dataData)
	collectionCommon := collectioncommon.NewCollectionCommon(collectionRepo)
	answerCommon := answercommon.NewAnswerCommon(answerRepo)
	metaRepo := meta.NewMetaRepo(dataData)
	metaService := meta2.NewMetaService(metaRepo)
	questionCommon := questioncommon.NewQuestionCommon(questionRepo, answerRepo, voteRepo, followRepo, tagCommonService, userCommon, collectionCommon, answerCommon, metaService, configService, activityQueueService, dataData)
	collectionService := service.NewCollectionService(collectionRepo, collectionGroupRepo, questionCommon)
	collectionController := controller.NewCollectionController(collectionService)
	answerActivityRepo := activity.NewAnswerActivityRepo(dataData, activityRepo, userRankRepo, notificationQueueService)
	answerActivityService := activity2.NewAnswerActivityService(answerActivityRepo, configService)
	questionService := service.NewQuestionService(questionRepo, tagCommonService, questionCommon, userCommon, userRepo, revisionService, metaService, collectionCommon, answerActivityService, emailService, notificationQueueService, activityQueueService, siteInfoCommonService)
	answerService := service.NewAnswerService(answerRepo, questionRepo, questionCommon, userCommon, collectionCommon, userRepo, revisionService, answerActivityService, answerCommon, voteRepo, emailService, userRoleRelService, notificationQueueService, activityQueueService)
	questionController := controller.NewQuestionController(questionService, answerService, rankService, siteInfoCommonService, captchaService)
	answerController := controller.NewAnswerController(answerService, rankService, captchaService)
	searchParser := search_parser.NewSearchParser(tagCommonService, userCommon)
	searchRepo := search_common.NewSearchRepo(dataData, uniqueIDRepo, userCommon)
	searchService := service.NewSearchService(searchParser, searchRepo)
	searchController := controller.NewSearchController(searchService, captchaService)
	serviceRevisionService := service.NewRevisionService(revisionRepo, userCommon, questionCommon, answerService, objService, questionRepo, answerRepo, tagRepo, tagCommonService, notificationQueueService, activityQueueService)
	revisionController := controller.NewRevisionController(serviceRevisionService, rankService)
	rankController := controller.NewRankController(rankService)
	reportHandle := report_handle_admin.NewReportHandle(questionCommon, commentRepo, configService, notificationQueueService)
	reportAdminService := report_admin.NewReportAdminService(reportRepo, userCommon, answerRepo, questionRepo, commentCommonRepo, reportHandle, configService, objService)
	controller_adminReportController := controller_admin.NewReportController(reportAdminService)
	userAdminRepo := user.NewUserAdminRepo(dataData, authRepo)
	userAdminService := user_admin.NewUserAdminService(userAdminRepo, userRoleRelService, authService, userCommon, userActiveActivityRepo, siteInfoCommonService, emailService)
	userAdminController := controller_admin.NewUserAdminController(userAdminService)
	reasonRepo := reason.NewReasonRepo(configService)
	reasonService := reason2.NewReasonService(reasonRepo)
	reasonController := controller.NewReasonController(reasonService)
	themeController := controller_admin.NewThemeController()
	siteInfoService := siteinfo.NewSiteInfoService(siteInfoRepo, siteInfoCommonService, emailService, tagCommonService, configService, questionCommon)
	siteInfoController := controller_admin.NewSiteInfoController(siteInfoService)
	controllerSiteInfoController := controller.NewSiteInfoController(siteInfoCommonService)
	notificationRepo := notification.NewNotificationRepo(dataData)
	notificationCommon := notificationcommon.NewNotificationCommon(dataData, notificationRepo, userCommon, activityRepo, followRepo, objService, notificationQueueService)
	notificationService := notification2.NewNotificationService(dataData, notificationRepo, notificationCommon, revisionService)
	notificationController := controller.NewNotificationController(notificationService, rankService)
	dashboardService := dashboard.NewDashboardService(questionRepo, answerRepo, commentCommonRepo, voteRepo, userRepo, reportRepo, configService, siteInfoCommonService, serviceConf, dataData)
	dashboardController := controller.NewDashboardController(dashboardService)
	uploaderService := uploader.NewUploaderService(serviceConf, siteInfoCommonService)
	uploadController := controller.NewUploadController(uploaderService)
	activityActivityRepo := activity.NewActivityRepo(dataData, configService)
	activityCommon := activity_common2.NewActivityCommon(activityRepo, activityQueueService)
	commentCommonService := comment_common.NewCommentCommonService(commentCommonRepo)
	activityService := activity2.NewActivityService(activityActivityRepo, userCommon, activityCommon, tagCommonService, objService, commentCommonService, revisionService, metaService, configService)
	activityController := controller.NewActivityController(activityService)
	roleController := controller_admin.NewRoleController(roleService)
	pluginConfigRepo := plugin_config.NewPluginConfigRepo(dataData)
	pluginCommonService := plugin_common.NewPluginCommonService(pluginConfigRepo, configService, dataData)
	pluginController := controller_admin.NewPluginController(pluginCommonService)
	permissionController := controller.NewPermissionController(rankService)
	answerAPIRouter := router.NewAnswerAPIRouter(langController, userController, commentController, reportController, voteController, tagController, followController, collectionController, questionController, answerController, searchController, revisionController, rankController, controller_adminReportController, userAdminController, reasonController, themeController, siteInfoController, controllerSiteInfoController, notificationController, dashboardController, uploadController, activityController, roleController, pluginController, permissionController)
	swaggerRouter := router.NewSwaggerRouter(swaggerConf)
	uiRouter := router.NewUIRouter(controllerSiteInfoController, siteInfoCommonService)
	authUserMiddleware := middleware.NewAuthUserMiddleware(authService, siteInfoCommonService)
	avatarMiddleware := middleware.NewAvatarMiddleware(serviceConf, uploaderService)
	shortIDMiddleware := middleware.NewShortIDMiddleware(siteInfoCommonService)
	templateRenderController := templaterender.NewTemplateRenderController(questionService, userService, tagService, answerService, commentService, siteInfoCommonService, questionRepo)
	templateController := controller.NewTemplateController(templateRenderController, siteInfoCommonService)
	templateRouter := router.NewTemplateRouter(templateController, templateRenderController, siteInfoController, authUserMiddleware)
	connectorController := controller.NewConnectorController(siteInfoCommonService, emailService, userExternalLoginService)
	userCenterLoginService := user_external_login2.NewUserCenterLoginService(userRepo, userCommon, userExternalLoginRepo, userActiveActivityRepo, siteInfoCommonService)
	userCenterController := controller.NewUserCenterController(userCenterLoginService, siteInfoCommonService)
	pluginAPIRouter := router.NewPluginAPIRouter(connectorController, userCenterController)
	ginEngine := server.NewHTTPServer(debug, staticRouter, answerAPIRouter, swaggerRouter, uiRouter, authUserMiddleware, avatarMiddleware, shortIDMiddleware, templateRouter, pluginAPIRouter)
	scheduledTaskManager := cron.NewScheduledTaskManager(siteInfoCommonService, questionService)
	application := newApplication(serverConf, ginEngine, scheduledTaskManager)
	return application, func() {
		cleanup2()
		cleanup()
	}, nil
}
