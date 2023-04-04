// Code generated by mockery v2.20.0. DO NOT EDIT.

package mocks

import (
	context "context"

	models "github.com/shellhub-io/shellhub/pkg/models"
	mock "github.com/stretchr/testify/mock"

	paginator "github.com/shellhub-io/shellhub/pkg/api/paginator"

	request "github.com/shellhub-io/shellhub/pkg/api/requests"

	response "github.com/shellhub-io/shellhub/pkg/api/responses"

	rsa "crypto/rsa"
)

// Service is an autogenerated mock type for the Service type
type Service struct {
	mock.Mock
}

// AddNamespaceUser provides a mock function with given fields: ctx, memberUsername, memberRole, tenantID, userID
func (_m *Service) AddNamespaceUser(ctx context.Context, memberUsername string, memberRole string, tenantID string, userID string) (*models.Namespace, error) {
	ret := _m.Called(ctx, memberUsername, memberRole, tenantID, userID)

	var r0 *models.Namespace
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string, string, string, string) (*models.Namespace, error)); ok {
		return rf(ctx, memberUsername, memberRole, tenantID, userID)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, string, string, string) *models.Namespace); ok {
		r0 = rf(ctx, memberUsername, memberRole, tenantID, userID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.Namespace)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, string, string, string) error); ok {
		r1 = rf(ctx, memberUsername, memberRole, tenantID, userID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// AddPublicKeyTag provides a mock function with given fields: ctx, tenant, fingerprint, tag
func (_m *Service) AddPublicKeyTag(ctx context.Context, tenant string, fingerprint string, tag string) error {
	ret := _m.Called(ctx, tenant, fingerprint, tag)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, string, string) error); ok {
		r0 = rf(ctx, tenant, fingerprint, tag)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// AuthCacheToken provides a mock function with given fields: ctx, tenant, id, token
func (_m *Service) AuthCacheToken(ctx context.Context, tenant string, id string, token string) error {
	ret := _m.Called(ctx, tenant, id, token)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, string, string) error); ok {
		r0 = rf(ctx, tenant, id, token)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// AuthDevice provides a mock function with given fields: ctx, req, remoteAddr
func (_m *Service) AuthDevice(ctx context.Context, req request.DeviceAuth, remoteAddr string) (*models.DeviceAuthResponse, error) {
	ret := _m.Called(ctx, req, remoteAddr)

	var r0 *models.DeviceAuthResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, request.DeviceAuth, string) (*models.DeviceAuthResponse, error)); ok {
		return rf(ctx, req, remoteAddr)
	}
	if rf, ok := ret.Get(0).(func(context.Context, request.DeviceAuth, string) *models.DeviceAuthResponse); ok {
		r0 = rf(ctx, req, remoteAddr)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.DeviceAuthResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, request.DeviceAuth, string) error); ok {
		r1 = rf(ctx, req, remoteAddr)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// AuthGetToken provides a mock function with given fields: ctx, tenant
func (_m *Service) AuthGetToken(ctx context.Context, tenant string) (*models.UserAuthResponse, error) {
	ret := _m.Called(ctx, tenant)

	var r0 *models.UserAuthResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (*models.UserAuthResponse, error)); ok {
		return rf(ctx, tenant)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) *models.UserAuthResponse); ok {
		r0 = rf(ctx, tenant)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.UserAuthResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, tenant)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// AuthIsCacheToken provides a mock function with given fields: ctx, tenant, id
func (_m *Service) AuthIsCacheToken(ctx context.Context, tenant string, id string) (bool, error) {
	ret := _m.Called(ctx, tenant, id)

	var r0 bool
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string, string) (bool, error)); ok {
		return rf(ctx, tenant, id)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, string) bool); ok {
		r0 = rf(ctx, tenant, id)
	} else {
		r0 = ret.Get(0).(bool)
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, string) error); ok {
		r1 = rf(ctx, tenant, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// AuthPublicKey provides a mock function with given fields: ctx, req
func (_m *Service) AuthPublicKey(ctx context.Context, req request.PublicKeyAuth) (*models.PublicKeyAuthResponse, error) {
	ret := _m.Called(ctx, req)

	var r0 *models.PublicKeyAuthResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, request.PublicKeyAuth) (*models.PublicKeyAuthResponse, error)); ok {
		return rf(ctx, req)
	}
	if rf, ok := ret.Get(0).(func(context.Context, request.PublicKeyAuth) *models.PublicKeyAuthResponse); ok {
		r0 = rf(ctx, req)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.PublicKeyAuthResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, request.PublicKeyAuth) error); ok {
		r1 = rf(ctx, req)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// AuthSwapToken provides a mock function with given fields: ctx, ID, tenant
func (_m *Service) AuthSwapToken(ctx context.Context, ID string, tenant string) (*models.UserAuthResponse, error) {
	ret := _m.Called(ctx, ID, tenant)

	var r0 *models.UserAuthResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string, string) (*models.UserAuthResponse, error)); ok {
		return rf(ctx, ID, tenant)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, string) *models.UserAuthResponse); ok {
		r0 = rf(ctx, ID, tenant)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.UserAuthResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, string) error); ok {
		r1 = rf(ctx, ID, tenant)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// AuthUncacheToken provides a mock function with given fields: ctx, tenant, id
func (_m *Service) AuthUncacheToken(ctx context.Context, tenant string, id string) error {
	ret := _m.Called(ctx, tenant, id)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, string) error); ok {
		r0 = rf(ctx, tenant, id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// AuthUser provides a mock function with given fields: ctx, req
func (_m *Service) AuthUser(ctx context.Context, req request.UserAuth) (*models.UserAuthResponse, error) {
	ret := _m.Called(ctx, req)

	var r0 *models.UserAuthResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, request.UserAuth) (*models.UserAuthResponse, error)); ok {
		return rf(ctx, req)
	}
	if rf, ok := ret.Get(0).(func(context.Context, request.UserAuth) *models.UserAuthResponse); ok {
		r0 = rf(ctx, req)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.UserAuthResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, request.UserAuth) error); ok {
		r1 = rf(ctx, req)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// AuthUserInfo provides a mock function with given fields: ctx, username, tenant, token
func (_m *Service) AuthUserInfo(ctx context.Context, username string, tenant string, token string) (*models.UserAuthResponse, error) {
	ret := _m.Called(ctx, username, tenant, token)

	var r0 *models.UserAuthResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string, string, string) (*models.UserAuthResponse, error)); ok {
		return rf(ctx, username, tenant, token)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, string, string) *models.UserAuthResponse); ok {
		r0 = rf(ctx, username, tenant, token)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.UserAuthResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, string, string) error); ok {
		r1 = rf(ctx, username, tenant, token)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CreateDeviceTag provides a mock function with given fields: ctx, uid, tag
func (_m *Service) CreateDeviceTag(ctx context.Context, uid models.UID, tag string) error {
	ret := _m.Called(ctx, uid, tag)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, models.UID, string) error); ok {
		r0 = rf(ctx, uid, tag)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// CreateNamespace provides a mock function with given fields: ctx, namespace, userID
func (_m *Service) CreateNamespace(ctx context.Context, namespace request.NamespaceCreate, userID string) (*models.Namespace, error) {
	ret := _m.Called(ctx, namespace, userID)

	var r0 *models.Namespace
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, request.NamespaceCreate, string) (*models.Namespace, error)); ok {
		return rf(ctx, namespace, userID)
	}
	if rf, ok := ret.Get(0).(func(context.Context, request.NamespaceCreate, string) *models.Namespace); ok {
		r0 = rf(ctx, namespace, userID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.Namespace)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, request.NamespaceCreate, string) error); ok {
		r1 = rf(ctx, namespace, userID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CreatePrivateKey provides a mock function with given fields: ctx
func (_m *Service) CreatePrivateKey(ctx context.Context) (*models.PrivateKey, error) {
	ret := _m.Called(ctx)

	var r0 *models.PrivateKey
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context) (*models.PrivateKey, error)); ok {
		return rf(ctx)
	}
	if rf, ok := ret.Get(0).(func(context.Context) *models.PrivateKey); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.PrivateKey)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CreatePublicKey provides a mock function with given fields: ctx, req, tenant
func (_m *Service) CreatePublicKey(ctx context.Context, req request.PublicKeyCreate, tenant string) (*response.PublicKeyCreate, error) {
	ret := _m.Called(ctx, req, tenant)

	var r0 *response.PublicKeyCreate
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, request.PublicKeyCreate, string) (*response.PublicKeyCreate, error)); ok {
		return rf(ctx, req, tenant)
	}
	if rf, ok := ret.Get(0).(func(context.Context, request.PublicKeyCreate, string) *response.PublicKeyCreate); ok {
		r0 = rf(ctx, req, tenant)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*response.PublicKeyCreate)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, request.PublicKeyCreate, string) error); ok {
		r1 = rf(ctx, req, tenant)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CreateSession provides a mock function with given fields: ctx, session
func (_m *Service) CreateSession(ctx context.Context, session request.SessionCreate) (*models.Session, error) {
	ret := _m.Called(ctx, session)

	var r0 *models.Session
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, request.SessionCreate) (*models.Session, error)); ok {
		return rf(ctx, session)
	}
	if rf, ok := ret.Get(0).(func(context.Context, request.SessionCreate) *models.Session); ok {
		r0 = rf(ctx, session)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.Session)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, request.SessionCreate) error); ok {
		r1 = rf(ctx, session)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeactivateSession provides a mock function with given fields: ctx, uid
func (_m *Service) DeactivateSession(ctx context.Context, uid models.UID) error {
	ret := _m.Called(ctx, uid)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, models.UID) error); ok {
		r0 = rf(ctx, uid)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DeleteDevice provides a mock function with given fields: ctx, uid, tenant
func (_m *Service) DeleteDevice(ctx context.Context, uid models.UID, tenant string) error {
	ret := _m.Called(ctx, uid, tenant)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, models.UID, string) error); ok {
		r0 = rf(ctx, uid, tenant)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DeleteNamespace provides a mock function with given fields: ctx, tenantID
func (_m *Service) DeleteNamespace(ctx context.Context, tenantID string) error {
	ret := _m.Called(ctx, tenantID)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string) error); ok {
		r0 = rf(ctx, tenantID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DeletePublicKey provides a mock function with given fields: ctx, fingerprint, tenant
func (_m *Service) DeletePublicKey(ctx context.Context, fingerprint string, tenant string) error {
	ret := _m.Called(ctx, fingerprint, tenant)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, string) error); ok {
		r0 = rf(ctx, fingerprint, tenant)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DeleteTag provides a mock function with given fields: ctx, tenant, tag
func (_m *Service) DeleteTag(ctx context.Context, tenant string, tag string) error {
	ret := _m.Called(ctx, tenant, tag)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, string) error); ok {
		r0 = rf(ctx, tenant, tag)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DeviceHeartbeat provides a mock function with given fields: ctx, uid
func (_m *Service) DeviceHeartbeat(ctx context.Context, uid models.UID) error {
	ret := _m.Called(ctx, uid)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, models.UID) error); ok {
		r0 = rf(ctx, uid)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// EditNamespace provides a mock function with given fields: ctx, tenantID, name
func (_m *Service) EditNamespace(ctx context.Context, tenantID string, name string) (*models.Namespace, error) {
	ret := _m.Called(ctx, tenantID, name)

	var r0 *models.Namespace
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string, string) (*models.Namespace, error)); ok {
		return rf(ctx, tenantID, name)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, string) *models.Namespace); ok {
		r0 = rf(ctx, tenantID, name)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.Namespace)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, string) error); ok {
		r1 = rf(ctx, tenantID, name)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// EditNamespaceUser provides a mock function with given fields: ctx, tenantID, userID, memberID, memberNewRole
func (_m *Service) EditNamespaceUser(ctx context.Context, tenantID string, userID string, memberID string, memberNewRole string) error {
	ret := _m.Called(ctx, tenantID, userID, memberID, memberNewRole)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, string, string, string) error); ok {
		r0 = rf(ctx, tenantID, userID, memberID, memberNewRole)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// EditSessionRecordStatus provides a mock function with given fields: ctx, sessionRecord, tenantID
func (_m *Service) EditSessionRecordStatus(ctx context.Context, sessionRecord bool, tenantID string) error {
	ret := _m.Called(ctx, sessionRecord, tenantID)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, bool, string) error); ok {
		r0 = rf(ctx, sessionRecord, tenantID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// EvaluateKeyFilter provides a mock function with given fields: ctx, key, dev
func (_m *Service) EvaluateKeyFilter(ctx context.Context, key *models.PublicKey, dev models.Device) (bool, error) {
	ret := _m.Called(ctx, key, dev)

	var r0 bool
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *models.PublicKey, models.Device) (bool, error)); ok {
		return rf(ctx, key, dev)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *models.PublicKey, models.Device) bool); ok {
		r0 = rf(ctx, key, dev)
	} else {
		r0 = ret.Get(0).(bool)
	}

	if rf, ok := ret.Get(1).(func(context.Context, *models.PublicKey, models.Device) error); ok {
		r1 = rf(ctx, key, dev)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// EvaluateKeyUsername provides a mock function with given fields: ctx, key, username
func (_m *Service) EvaluateKeyUsername(ctx context.Context, key *models.PublicKey, username string) (bool, error) {
	ret := _m.Called(ctx, key, username)

	var r0 bool
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *models.PublicKey, string) (bool, error)); ok {
		return rf(ctx, key, username)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *models.PublicKey, string) bool); ok {
		r0 = rf(ctx, key, username)
	} else {
		r0 = ret.Get(0).(bool)
	}

	if rf, ok := ret.Get(1).(func(context.Context, *models.PublicKey, string) error); ok {
		r1 = rf(ctx, key, username)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetDevice provides a mock function with given fields: ctx, uid
func (_m *Service) GetDevice(ctx context.Context, uid models.UID) (*models.Device, error) {
	ret := _m.Called(ctx, uid)

	var r0 *models.Device
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, models.UID) (*models.Device, error)); ok {
		return rf(ctx, uid)
	}
	if rf, ok := ret.Get(0).(func(context.Context, models.UID) *models.Device); ok {
		r0 = rf(ctx, uid)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.Device)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, models.UID) error); ok {
		r1 = rf(ctx, uid)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetNamespace provides a mock function with given fields: ctx, tenantID
func (_m *Service) GetNamespace(ctx context.Context, tenantID string) (*models.Namespace, error) {
	ret := _m.Called(ctx, tenantID)

	var r0 *models.Namespace
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (*models.Namespace, error)); ok {
		return rf(ctx, tenantID)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) *models.Namespace); ok {
		r0 = rf(ctx, tenantID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.Namespace)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, tenantID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetPublicKey provides a mock function with given fields: ctx, fingerprint, tenant
func (_m *Service) GetPublicKey(ctx context.Context, fingerprint string, tenant string) (*models.PublicKey, error) {
	ret := _m.Called(ctx, fingerprint, tenant)

	var r0 *models.PublicKey
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string, string) (*models.PublicKey, error)); ok {
		return rf(ctx, fingerprint, tenant)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, string) *models.PublicKey); ok {
		r0 = rf(ctx, fingerprint, tenant)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.PublicKey)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, string) error); ok {
		r1 = rf(ctx, fingerprint, tenant)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetSession provides a mock function with given fields: ctx, uid
func (_m *Service) GetSession(ctx context.Context, uid models.UID) (*models.Session, error) {
	ret := _m.Called(ctx, uid)

	var r0 *models.Session
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, models.UID) (*models.Session, error)); ok {
		return rf(ctx, uid)
	}
	if rf, ok := ret.Get(0).(func(context.Context, models.UID) *models.Session); ok {
		r0 = rf(ctx, uid)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.Session)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, models.UID) error); ok {
		r1 = rf(ctx, uid)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetSessionRecord provides a mock function with given fields: ctx, tenantID
func (_m *Service) GetSessionRecord(ctx context.Context, tenantID string) (bool, error) {
	ret := _m.Called(ctx, tenantID)

	var r0 bool
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (bool, error)); ok {
		return rf(ctx, tenantID)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) bool); ok {
		r0 = rf(ctx, tenantID)
	} else {
		r0 = ret.Get(0).(bool)
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, tenantID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetStats provides a mock function with given fields: ctx
func (_m *Service) GetStats(ctx context.Context) (*models.Stats, error) {
	ret := _m.Called(ctx)

	var r0 *models.Stats
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context) (*models.Stats, error)); ok {
		return rf(ctx)
	}
	if rf, ok := ret.Get(0).(func(context.Context) *models.Stats); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.Stats)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetTags provides a mock function with given fields: ctx, tenant
func (_m *Service) GetTags(ctx context.Context, tenant string) ([]string, int, error) {
	ret := _m.Called(ctx, tenant)

	var r0 []string
	var r1 int
	var r2 error
	if rf, ok := ret.Get(0).(func(context.Context, string) ([]string, int, error)); ok {
		return rf(ctx, tenant)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) []string); ok {
		r0 = rf(ctx, tenant)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]string)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) int); ok {
		r1 = rf(ctx, tenant)
	} else {
		r1 = ret.Get(1).(int)
	}

	if rf, ok := ret.Get(2).(func(context.Context, string) error); ok {
		r2 = rf(ctx, tenant)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// KeepAliveSession provides a mock function with given fields: ctx, uid
func (_m *Service) KeepAliveSession(ctx context.Context, uid models.UID) error {
	ret := _m.Called(ctx, uid)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, models.UID) error); ok {
		r0 = rf(ctx, uid)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// ListDevices provides a mock function with given fields: ctx, tenant, pagination, filter, status, sort, order
func (_m *Service) ListDevices(ctx context.Context, tenant string, pagination paginator.Query, filter []models.Filter, status models.DeviceStatus, sort string, order string) ([]models.Device, int, error) {
	ret := _m.Called(ctx, tenant, pagination, filter, status, sort, order)

	var r0 []models.Device
	var r1 int
	var r2 error
	if rf, ok := ret.Get(0).(func(context.Context, string, paginator.Query, []models.Filter, models.DeviceStatus, string, string) ([]models.Device, int, error)); ok {
		return rf(ctx, tenant, pagination, filter, status, sort, order)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, paginator.Query, []models.Filter, models.DeviceStatus, string, string) []models.Device); ok {
		r0 = rf(ctx, tenant, pagination, filter, status, sort, order)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]models.Device)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, paginator.Query, []models.Filter, models.DeviceStatus, string, string) int); ok {
		r1 = rf(ctx, tenant, pagination, filter, status, sort, order)
	} else {
		r1 = ret.Get(1).(int)
	}

	if rf, ok := ret.Get(2).(func(context.Context, string, paginator.Query, []models.Filter, models.DeviceStatus, string, string) error); ok {
		r2 = rf(ctx, tenant, pagination, filter, status, sort, order)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// ListNamespaces provides a mock function with given fields: ctx, pagination, filter, export
func (_m *Service) ListNamespaces(ctx context.Context, pagination paginator.Query, filter []models.Filter, export bool) ([]models.Namespace, int, error) {
	ret := _m.Called(ctx, pagination, filter, export)

	var r0 []models.Namespace
	var r1 int
	var r2 error
	if rf, ok := ret.Get(0).(func(context.Context, paginator.Query, []models.Filter, bool) ([]models.Namespace, int, error)); ok {
		return rf(ctx, pagination, filter, export)
	}
	if rf, ok := ret.Get(0).(func(context.Context, paginator.Query, []models.Filter, bool) []models.Namespace); ok {
		r0 = rf(ctx, pagination, filter, export)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]models.Namespace)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, paginator.Query, []models.Filter, bool) int); ok {
		r1 = rf(ctx, pagination, filter, export)
	} else {
		r1 = ret.Get(1).(int)
	}

	if rf, ok := ret.Get(2).(func(context.Context, paginator.Query, []models.Filter, bool) error); ok {
		r2 = rf(ctx, pagination, filter, export)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// ListPublicKeys provides a mock function with given fields: ctx, pagination
func (_m *Service) ListPublicKeys(ctx context.Context, pagination paginator.Query) ([]models.PublicKey, int, error) {
	ret := _m.Called(ctx, pagination)

	var r0 []models.PublicKey
	var r1 int
	var r2 error
	if rf, ok := ret.Get(0).(func(context.Context, paginator.Query) ([]models.PublicKey, int, error)); ok {
		return rf(ctx, pagination)
	}
	if rf, ok := ret.Get(0).(func(context.Context, paginator.Query) []models.PublicKey); ok {
		r0 = rf(ctx, pagination)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]models.PublicKey)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, paginator.Query) int); ok {
		r1 = rf(ctx, pagination)
	} else {
		r1 = ret.Get(1).(int)
	}

	if rf, ok := ret.Get(2).(func(context.Context, paginator.Query) error); ok {
		r2 = rf(ctx, pagination)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// ListSessions provides a mock function with given fields: ctx, pagination
func (_m *Service) ListSessions(ctx context.Context, pagination paginator.Query) ([]models.Session, int, error) {
	ret := _m.Called(ctx, pagination)

	var r0 []models.Session
	var r1 int
	var r2 error
	if rf, ok := ret.Get(0).(func(context.Context, paginator.Query) ([]models.Session, int, error)); ok {
		return rf(ctx, pagination)
	}
	if rf, ok := ret.Get(0).(func(context.Context, paginator.Query) []models.Session); ok {
		r0 = rf(ctx, pagination)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]models.Session)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, paginator.Query) int); ok {
		r1 = rf(ctx, pagination)
	} else {
		r1 = ret.Get(1).(int)
	}

	if rf, ok := ret.Get(2).(func(context.Context, paginator.Query) error); ok {
		r2 = rf(ctx, pagination)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// LookupDevice provides a mock function with given fields: ctx, namespace, name
func (_m *Service) LookupDevice(ctx context.Context, namespace string, name string) (*models.Device, error) {
	ret := _m.Called(ctx, namespace, name)

	var r0 *models.Device
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string, string) (*models.Device, error)); ok {
		return rf(ctx, namespace, name)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, string) *models.Device); ok {
		r0 = rf(ctx, namespace, name)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.Device)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, string) error); ok {
		r1 = rf(ctx, namespace, name)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// PublicKey provides a mock function with given fields:
func (_m *Service) PublicKey() *rsa.PublicKey {
	ret := _m.Called()

	var r0 *rsa.PublicKey
	if rf, ok := ret.Get(0).(func() *rsa.PublicKey); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*rsa.PublicKey)
		}
	}

	return r0
}

// RemoveDeviceTag provides a mock function with given fields: ctx, uid, tag
func (_m *Service) RemoveDeviceTag(ctx context.Context, uid models.UID, tag string) error {
	ret := _m.Called(ctx, uid, tag)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, models.UID, string) error); ok {
		r0 = rf(ctx, uid, tag)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// RemoveNamespaceUser provides a mock function with given fields: ctx, tenantID, memberID, userID
func (_m *Service) RemoveNamespaceUser(ctx context.Context, tenantID string, memberID string, userID string) (*models.Namespace, error) {
	ret := _m.Called(ctx, tenantID, memberID, userID)

	var r0 *models.Namespace
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string, string, string) (*models.Namespace, error)); ok {
		return rf(ctx, tenantID, memberID, userID)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, string, string) *models.Namespace); ok {
		r0 = rf(ctx, tenantID, memberID, userID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.Namespace)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, string, string) error); ok {
		r1 = rf(ctx, tenantID, memberID, userID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// RemovePublicKeyTag provides a mock function with given fields: ctx, tenant, fingerprint, tag
func (_m *Service) RemovePublicKeyTag(ctx context.Context, tenant string, fingerprint string, tag string) error {
	ret := _m.Called(ctx, tenant, fingerprint, tag)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, string, string) error); ok {
		r0 = rf(ctx, tenant, fingerprint, tag)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// RenameDevice provides a mock function with given fields: ctx, uid, name, tenant
func (_m *Service) RenameDevice(ctx context.Context, uid models.UID, name string, tenant string) error {
	ret := _m.Called(ctx, uid, name, tenant)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, models.UID, string, string) error); ok {
		r0 = rf(ctx, uid, name, tenant)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// RenameTag provides a mock function with given fields: ctx, tenant, oldTag, newTag
func (_m *Service) RenameTag(ctx context.Context, tenant string, oldTag string, newTag string) error {
	ret := _m.Called(ctx, tenant, oldTag, newTag)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, string, string) error); ok {
		r0 = rf(ctx, tenant, oldTag, newTag)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// SetDevicePosition provides a mock function with given fields: ctx, uid, ip
func (_m *Service) SetDevicePosition(ctx context.Context, uid models.UID, ip string) error {
	ret := _m.Called(ctx, uid, ip)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, models.UID, string) error); ok {
		r0 = rf(ctx, uid, ip)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// SetSessionAuthenticated provides a mock function with given fields: ctx, uid, authenticated
func (_m *Service) SetSessionAuthenticated(ctx context.Context, uid models.UID, authenticated bool) error {
	ret := _m.Called(ctx, uid, authenticated)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, models.UID, bool) error); ok {
		r0 = rf(ctx, uid, authenticated)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Setup provides a mock function with given fields: ctx, req
func (_m *Service) Setup(ctx context.Context, req request.Setup) error {
	ret := _m.Called(ctx, req)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, request.Setup) error); ok {
		r0 = rf(ctx, req)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// UpdateDataUser provides a mock function with given fields: ctx, id, userData
func (_m *Service) UpdateDataUser(ctx context.Context, id string, userData request.UserDataUpdate) ([]string, error) {
	ret := _m.Called(ctx, id, userData)

	var r0 []string
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string, request.UserDataUpdate) ([]string, error)); ok {
		return rf(ctx, id, userData)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, request.UserDataUpdate) []string); ok {
		r0 = rf(ctx, id, userData)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]string)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, request.UserDataUpdate) error); ok {
		r1 = rf(ctx, id, userData)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateDevice provides a mock function with given fields: ctx, tenant, uid, name, publicURL
func (_m *Service) UpdateDevice(ctx context.Context, tenant string, uid models.UID, name *string, publicURL *bool) error {
	ret := _m.Called(ctx, tenant, uid, name, publicURL)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, models.UID, *string, *bool) error); ok {
		r0 = rf(ctx, tenant, uid, name, publicURL)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// UpdateDeviceStatus provides a mock function with given fields: ctx, uid, online
func (_m *Service) UpdateDeviceStatus(ctx context.Context, uid models.UID, online bool) error {
	ret := _m.Called(ctx, uid, online)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, models.UID, bool) error); ok {
		r0 = rf(ctx, uid, online)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// UpdateDeviceTag provides a mock function with given fields: ctx, uid, tags
func (_m *Service) UpdateDeviceTag(ctx context.Context, uid models.UID, tags []string) error {
	ret := _m.Called(ctx, uid, tags)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, models.UID, []string) error); ok {
		r0 = rf(ctx, uid, tags)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// UpdatePasswordUser provides a mock function with given fields: ctx, id, currentPassword, newPassword
func (_m *Service) UpdatePasswordUser(ctx context.Context, id string, currentPassword string, newPassword string) error {
	ret := _m.Called(ctx, id, currentPassword, newPassword)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, string, string) error); ok {
		r0 = rf(ctx, id, currentPassword, newPassword)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// UpdatePendingStatus provides a mock function with given fields: ctx, uid, status, tenant
func (_m *Service) UpdatePendingStatus(ctx context.Context, uid models.UID, status models.DeviceStatus, tenant string) error {
	ret := _m.Called(ctx, uid, status, tenant)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, models.UID, models.DeviceStatus, string) error); ok {
		r0 = rf(ctx, uid, status, tenant)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// UpdatePublicKey provides a mock function with given fields: ctx, fingerprint, tenant, key
func (_m *Service) UpdatePublicKey(ctx context.Context, fingerprint string, tenant string, key request.PublicKeyUpdate) (*models.PublicKey, error) {
	ret := _m.Called(ctx, fingerprint, tenant, key)

	var r0 *models.PublicKey
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string, string, request.PublicKeyUpdate) (*models.PublicKey, error)); ok {
		return rf(ctx, fingerprint, tenant, key)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, string, request.PublicKeyUpdate) *models.PublicKey); ok {
		r0 = rf(ctx, fingerprint, tenant, key)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.PublicKey)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, string, request.PublicKeyUpdate) error); ok {
		r1 = rf(ctx, fingerprint, tenant, key)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdatePublicKeyTags provides a mock function with given fields: ctx, tenant, fingerprint, tags
func (_m *Service) UpdatePublicKeyTags(ctx context.Context, tenant string, fingerprint string, tags []string) error {
	ret := _m.Called(ctx, tenant, fingerprint, tags)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, string, []string) error); ok {
		r0 = rf(ctx, tenant, fingerprint, tags)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

type mockConstructorTestingTNewService interface {
	mock.TestingT
	Cleanup(func())
}

// NewService creates a new instance of Service. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewService(t mockConstructorTestingTNewService) *Service {
	mock := &Service{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
