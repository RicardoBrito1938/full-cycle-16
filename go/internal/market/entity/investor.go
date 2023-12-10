package entity

type Investor struct {
	Id            string
	Name          string
	AssetPosition []*InvestorPosition
}

func NewInvestor(id, name string) *Investor {
	return &Investor{
		Id:            id,
		AssetPosition: []*InvestorPosition{},
	}
}

func (i *Investor) AddAssetPosition(AssetPosition *InvestorPosition) {
	i.AssetPosition = append(i.AssetPosition, AssetPosition)
}

func (i *Investor) GetAssetPosition(assetId string) *InvestorPosition {
	for _, v := range i.AssetPosition {
		if v.AssetID == assetId {
			return v
		}
	}
	return nil
}

func NewInvestorAssetPosition(assetId string, shares int) *InvestorPosition {
	return &InvestorPosition{
		AssetID: assetId,
		Shares:  shares,
	}
}

func (i *Investor) UpdateAssetPosition(assetId string, shares int) {
	assetPosition := i.GetAssetPosition(assetId)
	if assetPosition == nil {
		i.AssetPosition = append(i.AssetPosition, NewInvestorAssetPosition(assetId, shares))
	} else {
		assetPosition.Shares += shares
	}

}

type InvestorPosition struct {
	AssetID string
	Shares  int
}
