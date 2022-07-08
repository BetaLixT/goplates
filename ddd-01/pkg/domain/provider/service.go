package provider

type ProviderService struct {

}

func (svc *ProviderService) CreateProvider(
  prov IServiceProvider,
  title string,
  desc string,
) (created Provider, err error) {
  created, err = prov.GetProviderRepo().Create(
    title,
    desc,
  )
  return
}

func (svc *ProviderService) GetProvider(
  prov IServiceProvider,
  id string, 
) (created Provider, err error) {
  created, err = prov.GetProviderRepo().Get(
    id,
  )
  return
}

func (svc *ProviderService) DeleteProvider(
  prov IServiceProvider,
  id string, 
) (created Provider, err error) {
  created, err = prov.GetProviderRepo().Delete(
    id,
  )
  return
}
