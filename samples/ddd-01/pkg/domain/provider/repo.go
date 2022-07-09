package provider

type IProviderRepo interface {
  Create(
    title string,
    desc string,
  ) (Provider, error)
  Get(id string) (Provider, error)
  Delete(id string) (Provider, error)
}
