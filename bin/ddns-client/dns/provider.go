package dns

import "masklr-ddns/config"

type Provider interface {
    Update(cfg *config.Config, ip string) error
}

func NewProvider(name string) Provider {
    switch name {
    case "noip":
        return &NoIPProvider{}
    default:
        panic("未知服务商: " + name)
    }
}