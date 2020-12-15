package types

import (
	"bytes"
	"errors"
	"fmt"
	"github.com/ethereum/go-ethereum/crypto"
	"math/big"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum/common"

	zeroex "github.com/InjectiveLabs/sdk-go"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/ethereum/go-ethereum/common/math"
)

const RouterKey = ModuleName

var (
	_ sdk.Msg = &MsgRegisterDerivativeMarket{}
	_ sdk.Msg = &MsgSuspendDerivativeMarket{}
	_ sdk.Msg = &MsgResumeDerivativeMarket{}
	_ sdk.Msg = &MsgRegisterDerivativeMarket{}
	_ sdk.Msg = &MsgCreateDerivativeOrder{}
	_ sdk.Msg = &MsgRegisterSpotMarket{}
	_ sdk.Msg = &MsgSuspendSpotMarket{}
	_ sdk.Msg = &MsgResumeSpotMarket{}
	_ sdk.Msg = &MsgCreateSpotOrder{}
)

// Route should return the name of the module
func (msg MsgCreateSpotOrder) Route() string { return RouterKey }

// Type should return the action
func (msg MsgCreateSpotOrder) Type() string { return "createSpotOrder" }

// ValidateBasic runs stateless checks on the message
func (msg MsgCreateSpotOrder) ValidateBasic() error {
	if msg.Sender == "" {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, msg.Sender)
	}

	quantity := BigNum(msg.Order.GetTakerAssetAmount()).Int()
	if msg.Order == nil {
		return sdkerrors.Wrap(sdkerrors.ErrUnknownRequest, "no make order specified")
	} else if _, err := msg.Order.ToSignedOrder().ComputeOrderHash(); err != nil {
		return sdkerrors.Wrap(sdkerrors.ErrUnknownRequest, fmt.Sprintf("hash check failed: %v", err))
	} else if !isValidSignature(msg.Order.Signature) {
		return sdkerrors.Wrap(sdkerrors.ErrUnknownRequest, "invalid signature")
	} else if quantity == nil || quantity.Cmp(big.NewInt(0)) <= 0 {
		return sdkerrors.Wrap(ErrInsufficientOrderQuantity, "insufficient quantity")
	}

	return nil
}

// isValidSignature checks that the signature of the order is correct
func isValidSignature(sig string) bool {
	signature := common.FromHex(sig)
	signatureType := zeroex.SignatureType(signature[len(signature)-1])

	switch signatureType {
	case zeroex.InvalidSignature, zeroex.IllegalSignature:
		return false

	case zeroex.EIP712Signature:
		if len(signature) != 66 {
			return false
		}
		// TODO: Do further validation by splitting into r,s,v and do ECRecover

	case zeroex.EthSignSignature:
		if len(signature) != 66 {
			return false
		}
		// TODO: Do further validation by splitting into r,s,v, add prefix to hash
		// and do ECRecover

	case zeroex.ValidatorSignature:
		if len(signature) < 21 {
			return false
		}

	case zeroex.PreSignedSignature, zeroex.WalletSignature, zeroex.EIP1271WalletSignature:
		return true

	default:
		return false
	}

	return true
}

// GetSignBytes encodes the message for signing
func (msg *MsgCreateSpotOrder) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(msg))
}

// GetSigners defines whose signature is required
func (msg MsgCreateSpotOrder) GetSigners() []sdk.AccAddress {
	sender, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{sender}
}

// Route should return the name of the module
func (msg MsgCreateDerivativeOrder) Route() string { return RouterKey }

// Type should return the action
func (msg MsgCreateDerivativeOrder) Type() string { return "createDerivativeOrder" }

// ValidateBasic runs stateless checks on the message
func (msg MsgCreateDerivativeOrder) ValidateBasic() error {
	if msg.Sender == "" {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, msg.Sender)
	}

	if msg.Order == nil {
		return sdkerrors.Wrap(sdkerrors.ErrUnknownRequest, "no make order specified")
	} else if _, err := msg.Order.ToSignedOrder().ComputeOrderHash(); err != nil {
		return sdkerrors.Wrap(sdkerrors.ErrUnknownRequest, fmt.Sprintf("hash check failed: %v", err))
	} else if !isValidSignature(msg.Order.Signature) {
		return sdkerrors.Wrap(sdkerrors.ErrUnknownRequest, "invalid signature")
	}

	return nil
}

// GetSignBytes encodes the message for signing
func (msg *MsgCreateDerivativeOrder) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(msg))
}

// GetSigners defines whose signature is required
func (msg MsgCreateDerivativeOrder) GetSigners() []sdk.AccAddress {
	sender, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{sender}
}

// Route should return the name of the module
func (msg MsgRegisterSpotMarket) Route() string { return RouterKey }

// Type should return the action
func (msg MsgRegisterSpotMarket) Type() string { return "registerSpotMarket" }

// ValidateBasic runs stateless checks on the message
func (msg MsgRegisterSpotMarket) ValidateBasic() error {
	if msg.Sender == "" {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, msg.Sender)
	}
	if len(msg.Name) == 0 {
		return sdkerrors.Wrap(ErrBadField, "no trade pair name specified")
	} else if parts := strings.Split(msg.Name, "/"); len(parts) != 2 ||
		len(strings.TrimSpace(parts[0])) == 0 || len(strings.TrimSpace(parts[1])) == 0 {
		return sdkerrors.Wrap(ErrBadField, "pair name must be in format AAA/BBB")
	}
	if len(msg.MakerAssetData) == 0 {
		return sdkerrors.Wrap(ErrBadField, "no maker asset data specified")
	} else if len(msg.TakerAssetData) == 0 {
		return sdkerrors.Wrap(ErrBadField, "no taker asset data specified")
	}
	return nil
}

// GetSignBytes encodes the message for signing
func (msg *MsgRegisterSpotMarket) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(msg))
}

// GetSigners defines whose signature is required
func (msg MsgRegisterSpotMarket) GetSigners() []sdk.AccAddress {
	sender, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{sender}
}

// Route should return the name of the module
func (msg MsgRegisterDerivativeMarket) Route() string { return RouterKey }

// Type should return the action
func (msg MsgRegisterDerivativeMarket) Type() string { return "registerDerivativeMarket" }

// ValidateBasic runs stateless checks on the message
func (msg MsgRegisterDerivativeMarket) ValidateBasic() error {
	if msg.Sender == "" {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, msg.Sender)
	}
	market := msg.Market

	if len(market.Ticker) == 0 {
		return sdkerrors.Wrap(ErrBadField, "no market ticker name specified")
	} else if parts := strings.Split(market.Ticker, "/"); len(parts) != 2 ||
		len(strings.TrimSpace(parts[0])) == 0 || len(strings.TrimSpace(parts[1])) == 0 {
		return sdkerrors.Wrap(ErrBadField, "market ticker must be in format AAA/BBB")
	}
	// TODO: (albertchon) proper validation here
	//hash, err := market.Hash()
	//if err != nil {
	//	return sdkerrors.Wrap(ErrMarketInvalid, err.Error())
	//}
	//if hash.String() != market.MarketId {
	//	errMsg := "The MarketID " + market.MarketId + " provided does not match the MarketID " + hash.String() + " computed"
	//	errMsg += "\n for Ticker: " + market.Ticker + "\nOracle: " + market.Oracle + "\nBaseCurrency: " + market.BaseCurrency + "\nNonce: " + market.Nonce
	//	return sdkerrors.Wrap(ErrMarketInvalid, errMsg)
	//}

	return nil
}

// GetSignBytes encodes the message for signing
func (msg *MsgRegisterDerivativeMarket) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(msg))
}

// GetSigners defines whose signature is required
func (msg MsgRegisterDerivativeMarket) GetSigners() []sdk.AccAddress {
	sender, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{sender}
}

// Route should return the name of the module
func (msg MsgSuspendDerivativeMarket) Route() string { return RouterKey }

// Type should return the action
func (msg MsgSuspendDerivativeMarket) Type() string {
	return "suspendDerivativeMarket"
}

// ValidateBasic runs stateless checks on the message
func (msg MsgSuspendDerivativeMarket) ValidateBasic() error {
	// TODO: albertchon proper length checks here
	if msg.Sender == "" {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, msg.Sender)
	} else if msg.MarketId == "" {
		return sdkerrors.Wrap(ErrBadField, "no derivative market ID specified")
	} else if msg.ExchangeAddress == "" {
		return sdkerrors.Wrap(ErrBadField, "no derivative exchange address specified")

	}
	return nil
}

// GetSignBytes encodes the message for signing
func (msg *MsgSuspendDerivativeMarket) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(msg))
}

// GetSigners defines whose signature is required
func (msg MsgSuspendDerivativeMarket) GetSigners() []sdk.AccAddress {
	sender, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{sender}
}

// Route should return the name of the module
func (msg MsgResumeDerivativeMarket) Route() string { return RouterKey }

// Type should return the action
func (msg MsgResumeDerivativeMarket) Type() string {
	return "resumeDerivativeMarket"
}

// ValidateBasic runs stateless checks on the message
func (msg MsgResumeDerivativeMarket) ValidateBasic() error {
	// TODO: albertchon proper validation
	if msg.Sender == "" {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, msg.Sender)
	} else if msg.MarketId == "" {
		return sdkerrors.Wrap(ErrBadField, "no derivative market ID specified")
	} else if msg.ExchangeAddress == "" {
		return sdkerrors.Wrap(ErrBadField, "no derivative market ID specified")
	}
	return nil
}

// GetSignBytes encodes the message for signing
func (msg *MsgResumeDerivativeMarket) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(msg))
}

// GetSigners defines whose signature is required
func (msg MsgResumeDerivativeMarket) GetSigners() []sdk.AccAddress {
	sender, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{sender}
}

// Route should return the name of the module
func (msg MsgSuspendSpotMarket) Route() string { return RouterKey }

// Type should return the action
func (msg MsgSuspendSpotMarket) Type() string { return "suspendSpotMarket" }

// ValidateBasic runs stateless checks on the message
func (msg MsgSuspendSpotMarket) ValidateBasic() error {
	if msg.Sender == "" {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, msg.Sender)
	}

	if len(msg.Name) == 0 {
		return sdkerrors.Wrap(ErrBadField, "no trade pair name specified")
	}

	return nil
}

// GetSignBytes encodes the message for signing
func (msg *MsgSuspendSpotMarket) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(msg))
}

// GetSigners defines whose signature is required
func (msg MsgSuspendSpotMarket) GetSigners() []sdk.AccAddress {
	sender, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{sender}
}

// Route should return the name of the module
func (msg MsgResumeSpotMarket) Route() string { return RouterKey }

// Type should return the action
func (msg MsgResumeSpotMarket) Type() string { return "resumeSpotMarket" }

// ValidateBasic runs stateless checks on the message
func (msg MsgResumeSpotMarket) ValidateBasic() error {
	if msg.Sender == "" {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, msg.Sender)
	}

	if len(msg.Name) == 0 {
		return sdkerrors.Wrap(ErrBadField, "no trade pair name specified")
	}

	return nil
}

// GetSignBytes encodes the message for signing
func (msg *MsgResumeSpotMarket) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(msg))
}

// GetSigners defines whose signature is required
func (msg MsgResumeSpotMarket) GetSigners() []sdk.AccAddress {
	sender, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{sender}
}

// SafeSignedOrder is a special signed order structure
// for including in Msgs, because it consists of primitive types.
// Avoid using raw *big.Int in Msgs.
type SafeSignedOrder struct {
	// ChainID is a network identifier of the order.
	ChainID int64 `json:"chainID,omitempty"`
	// Exchange v3 contract address.
	ExchangeAddress Address `json:"exchangeAddress,omitempty"`
	// Address that created the order.
	MakerAddress Address `json:"makerAddress,omitempty"`
	// Address that is allowed to fill the order. If set to "0x0", any address is
	// allowed to fill the order.
	TakerAddress Address `json:"takerAddress,omitempty"`
	// Address that will receive fees when order is filled.
	FeeRecipientAddress Address `json:"feeRecipientAddress,omitempty"`
	// Address that is allowed to call Exchange contract methods that affect this
	// order. If set to "0x0", any address is allowed to call these methods.
	SenderAddress Address `json:"senderAddress,omitempty"`
	// Amount of makerAsset being offered by maker. Must be greater than 0.
	MakerAssetAmount BigNum `json:"makerAssetAmount,omitempty"`
	// Amount of takerAsset being bid on by maker. Must be greater than 0.
	TakerAssetAmount BigNum `json:"takerAssetAmount,omitempty"`
	// Amount of Fee Asset paid to feeRecipientAddress by maker when order is filled. If set to
	// 0, no transfer of Fee Asset from maker to feeRecipientAddress will be attempted.
	MakerFee BigNum `json:"makerFee,omitempty"`
	// Amount of Fee Asset paid to feeRecipientAddress by taker when order is filled. If set to
	// 0, no transfer of Fee Asset from taker to feeRecipientAddress will be attempted.
	TakerFee BigNum `json:"takerFee,omitempty"`
	// Timestamp in seconds at which order expires.
	ExpirationTimeSeconds BigNum `json:"expirationTimeSeconds,omitempty"`
	// Arbitrary number to facilitate uniqueness of the order's hash.
	Salt BigNum `json:"salt,omitempty"`
	// ABIv2 encoded data that can be decoded by a specified proxy contract when
	// transferring makerAsset.
	MakerAssetData HexBytes `json:"makerAssetData,omitempty"`
	// ABIv2 encoded data that can be decoded by a specified proxy contract when
	// transferring takerAsset.
	TakerAssetData HexBytes `json:"takerAssetData,omitempty"`
	// ABIv2 encoded data that can be decoded by a specified proxy contract when
	// transferring makerFee.
	MakerFeeAssetData HexBytes `json:"makerFeeAssetData,omitempty"`
	// ABIv2 encoded data that can be decoded by a specified proxy contract when
	// transferring takerFee.
	TakerFeeAssetData HexBytes `json:"takerFeeAssetData,omitempty"`
	// Order signature.
	Signature HexBytes `json:"signature,omitempty"`
}

// NewSafeSignedOrder constructs a new SafeSignedOrder from given zeroex.SignedOrder.
func NewSafeSignedOrder(o *zeroex.SignedOrder) *SafeSignedOrder {
	return zo2so(o)
}

// ToSignedOrder returns an appropriate zeroex.SignedOrder defined by SafeSignedOrder.
func (m *BaseOrder) ToSignedOrder() *zeroex.SignedOrder {
	o, err := so2zo(m)
	if err != nil {
		panic(err)
	}
	return o
}

func (order *Order) DoesValidationPass(isLong bool, market *DerivativeMarket, currBlockTime time.Time) error {
	err := order.ComputeAndSetOrderType()
	if err != nil {
		return err
	}
	orderExpirationTime := BigNum(order.GetOrder().GetExpirationTimeSeconds()).Int()
	blockTime := big.NewInt(currBlockTime.Unix())

	if orderExpirationTime.Cmp(blockTime) <= 0 {
		return sdkerrors.Wrapf(ErrOrderExpired, "order expiration %s <= block time %s", orderExpirationTime.String(), blockTime.String())
	}
	margin := BigNum(order.Order.GetMakerFee()).Int()
	contractPriceMarginRequirement := order.ComputeContractPriceMarginRequirement(market)
	if margin.Cmp(contractPriceMarginRequirement) < 0 {
		return sdkerrors.Wrapf(ErrOverLeveragedOrder, "margin %s < contractPriceMarginRequirement %s", margin.String(), contractPriceMarginRequirement.String())
	}

	indexPriceMarginRequirement := order.ComputeIndexPriceMarginRequirement(isLong, market)
	indexPrice := BigNum(market.GetIndexPrice()).Int()

	if isLong && indexPrice.Cmp(indexPriceMarginRequirement) < 0 {
		return sdkerrors.Wrapf(ErrOverLeveragedOrder, "indexPrice %s <= indexPriceReq %s", market.GetIndexPrice(), order.IndexPriceRequirement)
	} else if !isLong && indexPrice.Cmp(indexPriceMarginRequirement) > 0 {
		return sdkerrors.Wrapf(ErrOverLeveragedOrder, "indexPrice %s >= indexPriceReq %s", market.GetIndexPrice(), order.IndexPriceRequirement)
	}
	return nil
}

func (order *Order) ComputeAndSetOrderType() error {
	orderTypeNumber := new(big.Int).SetBytes(common.FromHex(order.GetOrder().GetMakerFeeAssetData())[:common.HashLength]).Uint64()
	if orderTypeNumber != 0 && orderTypeNumber != 5 {
		return sdkerrors.Wrapf(ErrUnrecognizedOrderType, "Cannot recognize MakerFeeAssetData of %s", order.GetOrder().GetMakerFeeAssetData())
	}
	order.OrderType = orderTypeNumber
	return nil
}

func (order *Order) ComputeIndexPriceMarginRequirement(isLong bool, market *DerivativeMarket) *big.Int {
	price := BigNum(order.Order.GetMakerAssetAmount()).Int()
	quantity := BigNum(order.Order.GetTakerAssetAmount()).Int()
	margin := BigNum(order.Order.GetMakerFee()).Int()
	pq := new(big.Int).Mul(price, quantity)
	alphaQuantity := ScalePermyriad(quantity, BigNum(market.InitialMarginRatio).Int())
	num := new(big.Int)
	denom := new(big.Int)

	if isLong {
		num = num.Sub(margin, pq)
		denom = denom.Sub(alphaQuantity, quantity)
	} else {
		num = num.Add(margin, pq)
		denom = denom.Add(alphaQuantity, quantity)
	}

	indexPriceReq := new(big.Int).Div(num, denom)
	order.IndexPriceRequirement = indexPriceReq.String()
	return indexPriceReq
}

// quantity * initialMarginRatio * price
func (order *Order) ComputeContractPriceMarginRequirement(market *DerivativeMarket) *big.Int {
	price := BigNum(order.Order.GetMakerAssetAmount()).Int()
	quantity := BigNum(order.Order.GetTakerAssetAmount()).Int()
	alphaQuantity := ScalePermyriad(quantity, BigNum(market.InitialMarginRatio).Int())
	return new(big.Int).Mul(alphaQuantity, price)
}

// orderMarginHold = 1.0015 * margin * (remainingQuantity) / order.quantity
func (o *Order) ComputeOrderMarginHold(remainingQuantity, makerTxFeePermyriad *big.Int) (orderMarginHold *big.Int) {
	margin := BigNum(o.GetOrder().GetMakerFee()).Int()
	scaledMargin := IncrementByScaledPermyriad(margin, makerTxFeePermyriad)
	originalQuantity := BigNum(o.GetOrder().GetTakerAssetAmount()).Int()

	// TODO: filledAmount should always be zero with TEC since there will be no UnknownOrderHash
	numerator := new(big.Int).Mul(scaledMargin, remainingQuantity)
	orderMarginHold = new(big.Int).Div(numerator, originalQuantity)
	return orderMarginHold
}

// return amount * (1 + permyriad/10000) = (amount + amount * permyriad/10000)
func IncrementByScaledPermyriad(amount, permyriad *big.Int) *big.Int {
	return new(big.Int).Add(amount, ScalePermyriad(amount, permyriad))
}

// return (amount * permyriad) / 10000
func ScalePermyriad(amount, permyriad *big.Int) *big.Int {
	PERMYRIAD_BASE := BigNum("10000").Int()
	scaleFactor := new(big.Int).Mul(amount, permyriad)
	return new(big.Int).Div(scaleFactor, PERMYRIAD_BASE)
}

func ComputeSubaccountID(makerAddress string, takerFee string) common.Hash {
	subaccountID := crypto.Keccak256Hash(
		common.HexToAddress(makerAddress).Bytes(),
		common.LeftPadBytes(BigNum(takerFee).Int().Bytes(), 32),
	)
	//suplog.Debugf("%s + %s, => %s", makerAddress, takerFee, subaccountID.Hex())
	return subaccountID
}

// GetDirectionMarketAndSubaccountID
func (m *BaseOrder) GetDirectionMarketAndSubaccountID() (isLong bool, marketID common.Hash, subaccountID common.Hash) {
	mData, tData := common.FromHex(m.GetMakerAssetData()), common.FromHex(m.GetTakerAssetData())

	if len(mData) > common.HashLength {
		mData = mData[:common.HashLength]
	}

	if len(tData) > common.HashLength {
		tData = tData[:common.HashLength]
	}

	if bytes.Equal(tData, common.Hash{}.Bytes()) {
		isLong = true
		marketID = common.BytesToHash(mData)
	} else {
		isLong = false
		marketID = common.BytesToHash(tData)
	}
	subaccountID = ComputeSubaccountID(m.GetMakerAddress(), m.GetTakerFee())

	return isLong, marketID, subaccountID
}

// zo2so internal function converts model from *zeroex.SignedOrder to *SafeSignedOrder.
func zo2so(o *zeroex.SignedOrder) *SafeSignedOrder {
	if o == nil {
		return nil
	}
	return &SafeSignedOrder{
		ChainID:               o.ChainID.Int64(),
		ExchangeAddress:       Address{o.ExchangeAddress},
		MakerAddress:          Address{o.MakerAddress},
		TakerAddress:          Address{o.TakerAddress},
		FeeRecipientAddress:   Address{o.FeeRecipientAddress},
		SenderAddress:         Address{o.SenderAddress},
		MakerAssetAmount:      BigNum(o.MakerAssetAmount.String()),
		TakerAssetAmount:      BigNum(o.TakerAssetAmount.String()),
		MakerFee:              BigNum(o.MakerFee.String()),
		TakerFee:              BigNum(o.TakerFee.String()),
		ExpirationTimeSeconds: BigNum(o.ExpirationTimeSeconds.String()),
		Salt:                  BigNum(o.Salt.String()),
		MakerAssetData:        o.MakerAssetData,
		TakerAssetData:        o.TakerAssetData,
		MakerFeeAssetData:     o.MakerFeeAssetData,
		TakerFeeAssetData:     o.TakerFeeAssetData,
		Signature:             o.Signature,
	}
}

// so2zo internal function converts model from *SafeSignedOrder to *zeroex.SignedOrder.
func so2zo(o *BaseOrder) (*zeroex.SignedOrder, error) {
	if o == nil {
		return nil, nil
	}
	order := zeroex.Order{
		ChainID:             big.NewInt(o.ChainId),
		ExchangeAddress:     common.HexToAddress(o.ExchangeAddress),
		MakerAddress:        common.HexToAddress(o.MakerAddress),
		TakerAddress:        common.HexToAddress(o.TakerAddress),
		SenderAddress:       common.HexToAddress(o.SenderAddress),
		FeeRecipientAddress: common.HexToAddress(o.FeeRecipientAddress),
		MakerAssetData:      common.FromHex(o.MakerAssetData),
		MakerFeeAssetData:   common.FromHex(o.MakerFeeAssetData),
		TakerAssetData:      common.FromHex(o.TakerAssetData),
		TakerFeeAssetData:   common.FromHex(o.TakerFeeAssetData),
	}

	if v, ok := math.ParseBig256(string(o.MakerAssetAmount)); !ok {
		return nil, errors.New("makerAssetAmmount parse failed")
	} else {
		order.MakerAssetAmount = v
	}
	if v, ok := math.ParseBig256(string(o.MakerFee)); !ok {
		return nil, errors.New("makerFee parse failed")
	} else {
		order.MakerFee = v
	}
	if v, ok := math.ParseBig256(string(o.TakerAssetAmount)); !ok {
		return nil, errors.New("takerAssetAmmount parse failed")
	} else {
		order.TakerAssetAmount = v
	}
	if v, ok := math.ParseBig256(string(o.TakerFee)); !ok {
		return nil, errors.New("takerFee parse failed")
	} else {
		order.TakerFee = v
	}
	if v, ok := math.ParseBig256(string(o.ExpirationTimeSeconds)); !ok {
		return nil, errors.New("expirationTimeSeconds parse failed")
	} else {
		order.ExpirationTimeSeconds = v
	}
	if v, ok := math.ParseBig256(string(o.Salt)); !ok {
		return nil, errors.New("salt parse failed")
	} else {
		order.Salt = v
	}
	signedOrder := &zeroex.SignedOrder{
		Order:     order,
		Signature: common.FromHex(o.Signature),
	}
	return signedOrder, nil
}
