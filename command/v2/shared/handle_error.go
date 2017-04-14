package shared

import (
	"code.cloudfoundry.org/cli/actor/sharedaction"
	"code.cloudfoundry.org/cli/actor/v2action"
	"code.cloudfoundry.org/cli/api/cloudcontroller/ccerror"
	"code.cloudfoundry.org/cli/api/uaa"
	"code.cloudfoundry.org/cli/command"
)

func HandleError(err error) error {
	switch e := err.(type) {
	case ccerror.APINotFoundError:
		return command.APINotFoundError{URL: e.URL}
	case ccerror.RequestError:
		return command.APIRequestError{Err: e.Err}
	case ccerror.SSLValidationHostnameError:
		return command.SSLCertErrorError{Message: e.Message}
	case ccerror.UnverifiedServerError:
		return command.InvalidSSLCertError{API: e.URL}

	case ccerror.JobFailedError:
		return JobFailedError{JobGUID: e.JobGUID}
	case ccerror.JobTimeoutError:
		return JobTimeoutError{JobGUID: e.JobGUID}

	case uaa.InvalidAuthTokenError:
		return InvalidRefreshTokenError{}

	case sharedaction.NotLoggedInError:
		return command.NotLoggedInError{BinaryName: e.BinaryName}
	case sharedaction.NoTargetedOrganizationError:
		return command.NoTargetedOrganizationError{BinaryName: e.BinaryName}
	case sharedaction.NoTargetedSpaceError:
		return command.NoTargetedSpaceError{BinaryName: e.BinaryName}
	case sharedaction.PluginNotFoundError:
		return PluginNotFoundError{Name: e.Name}

	case v2action.ApplicationNotFoundError:
		return command.ApplicationNotFoundError{Name: e.Name}
	case v2action.OrganizationNotFoundError:
		return OrganizationNotFoundError{Name: e.Name}
	case v2action.SecurityGroupNotFoundError:
		return SecurityGroupNotFoundError{Name: e.Name}
	case v2action.ServiceInstanceNotFoundError:
		return command.ServiceInstanceNotFoundError{Name: e.Name}
	case v2action.SpaceNotFoundError:
		return SpaceNotFoundError{Name: e.Name}
	case v2action.HTTPHealthCheckInvalidError:
		return HTTPHealthCheckInvalidError{}
	}

	return err
}
