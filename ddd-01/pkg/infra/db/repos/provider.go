package repos

import "ddd-01/pkg/domain/provider"



type ProviderRepo struct {

}

var _ provider.IProviderRepo = (*ProviderRepo)(nil)
