//Package newordermultileg msg type = AB.
package newordermultileg

import (
	"github.com/quickfixgo/quickfix"
	"github.com/quickfixgo/quickfix/enum"
	"github.com/quickfixgo/quickfix/field"
)

//Message is a NewOrderMultileg wrapper for the generic Message type
type Message struct {
	quickfix.Message
}

//ClOrdID is a required field for NewOrderMultileg.
func (m Message) ClOrdID() (*field.ClOrdIDField, quickfix.MessageRejectError) {
	f := &field.ClOrdIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetClOrdID reads a ClOrdID from NewOrderMultileg.
func (m Message) GetClOrdID(f *field.ClOrdIDField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//SecondaryClOrdID is a non-required field for NewOrderMultileg.
func (m Message) SecondaryClOrdID() (*field.SecondaryClOrdIDField, quickfix.MessageRejectError) {
	f := &field.SecondaryClOrdIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSecondaryClOrdID reads a SecondaryClOrdID from NewOrderMultileg.
func (m Message) GetSecondaryClOrdID(f *field.SecondaryClOrdIDField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//ClOrdLinkID is a non-required field for NewOrderMultileg.
func (m Message) ClOrdLinkID() (*field.ClOrdLinkIDField, quickfix.MessageRejectError) {
	f := &field.ClOrdLinkIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetClOrdLinkID reads a ClOrdLinkID from NewOrderMultileg.
func (m Message) GetClOrdLinkID(f *field.ClOrdLinkIDField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//NoPartyIDs is a non-required field for NewOrderMultileg.
func (m Message) NoPartyIDs() (*field.NoPartyIDsField, quickfix.MessageRejectError) {
	f := &field.NoPartyIDsField{}
	err := m.Body.Get(f)
	return f, err
}

//GetNoPartyIDs reads a NoPartyIDs from NewOrderMultileg.
func (m Message) GetNoPartyIDs(f *field.NoPartyIDsField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//TradeOriginationDate is a non-required field for NewOrderMultileg.
func (m Message) TradeOriginationDate() (*field.TradeOriginationDateField, quickfix.MessageRejectError) {
	f := &field.TradeOriginationDateField{}
	err := m.Body.Get(f)
	return f, err
}

//GetTradeOriginationDate reads a TradeOriginationDate from NewOrderMultileg.
func (m Message) GetTradeOriginationDate(f *field.TradeOriginationDateField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//TradeDate is a non-required field for NewOrderMultileg.
func (m Message) TradeDate() (*field.TradeDateField, quickfix.MessageRejectError) {
	f := &field.TradeDateField{}
	err := m.Body.Get(f)
	return f, err
}

//GetTradeDate reads a TradeDate from NewOrderMultileg.
func (m Message) GetTradeDate(f *field.TradeDateField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//Account is a non-required field for NewOrderMultileg.
func (m Message) Account() (*field.AccountField, quickfix.MessageRejectError) {
	f := &field.AccountField{}
	err := m.Body.Get(f)
	return f, err
}

//GetAccount reads a Account from NewOrderMultileg.
func (m Message) GetAccount(f *field.AccountField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//AcctIDSource is a non-required field for NewOrderMultileg.
func (m Message) AcctIDSource() (*field.AcctIDSourceField, quickfix.MessageRejectError) {
	f := &field.AcctIDSourceField{}
	err := m.Body.Get(f)
	return f, err
}

//GetAcctIDSource reads a AcctIDSource from NewOrderMultileg.
func (m Message) GetAcctIDSource(f *field.AcctIDSourceField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//AccountType is a non-required field for NewOrderMultileg.
func (m Message) AccountType() (*field.AccountTypeField, quickfix.MessageRejectError) {
	f := &field.AccountTypeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetAccountType reads a AccountType from NewOrderMultileg.
func (m Message) GetAccountType(f *field.AccountTypeField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//DayBookingInst is a non-required field for NewOrderMultileg.
func (m Message) DayBookingInst() (*field.DayBookingInstField, quickfix.MessageRejectError) {
	f := &field.DayBookingInstField{}
	err := m.Body.Get(f)
	return f, err
}

//GetDayBookingInst reads a DayBookingInst from NewOrderMultileg.
func (m Message) GetDayBookingInst(f *field.DayBookingInstField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//BookingUnit is a non-required field for NewOrderMultileg.
func (m Message) BookingUnit() (*field.BookingUnitField, quickfix.MessageRejectError) {
	f := &field.BookingUnitField{}
	err := m.Body.Get(f)
	return f, err
}

//GetBookingUnit reads a BookingUnit from NewOrderMultileg.
func (m Message) GetBookingUnit(f *field.BookingUnitField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//PreallocMethod is a non-required field for NewOrderMultileg.
func (m Message) PreallocMethod() (*field.PreallocMethodField, quickfix.MessageRejectError) {
	f := &field.PreallocMethodField{}
	err := m.Body.Get(f)
	return f, err
}

//GetPreallocMethod reads a PreallocMethod from NewOrderMultileg.
func (m Message) GetPreallocMethod(f *field.PreallocMethodField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//AllocID is a non-required field for NewOrderMultileg.
func (m Message) AllocID() (*field.AllocIDField, quickfix.MessageRejectError) {
	f := &field.AllocIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetAllocID reads a AllocID from NewOrderMultileg.
func (m Message) GetAllocID(f *field.AllocIDField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//NoAllocs is a non-required field for NewOrderMultileg.
func (m Message) NoAllocs() (*field.NoAllocsField, quickfix.MessageRejectError) {
	f := &field.NoAllocsField{}
	err := m.Body.Get(f)
	return f, err
}

//GetNoAllocs reads a NoAllocs from NewOrderMultileg.
func (m Message) GetNoAllocs(f *field.NoAllocsField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//SettlType is a non-required field for NewOrderMultileg.
func (m Message) SettlType() (*field.SettlTypeField, quickfix.MessageRejectError) {
	f := &field.SettlTypeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSettlType reads a SettlType from NewOrderMultileg.
func (m Message) GetSettlType(f *field.SettlTypeField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//SettlDate is a non-required field for NewOrderMultileg.
func (m Message) SettlDate() (*field.SettlDateField, quickfix.MessageRejectError) {
	f := &field.SettlDateField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSettlDate reads a SettlDate from NewOrderMultileg.
func (m Message) GetSettlDate(f *field.SettlDateField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//CashMargin is a non-required field for NewOrderMultileg.
func (m Message) CashMargin() (*field.CashMarginField, quickfix.MessageRejectError) {
	f := &field.CashMarginField{}
	err := m.Body.Get(f)
	return f, err
}

//GetCashMargin reads a CashMargin from NewOrderMultileg.
func (m Message) GetCashMargin(f *field.CashMarginField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//ClearingFeeIndicator is a non-required field for NewOrderMultileg.
func (m Message) ClearingFeeIndicator() (*field.ClearingFeeIndicatorField, quickfix.MessageRejectError) {
	f := &field.ClearingFeeIndicatorField{}
	err := m.Body.Get(f)
	return f, err
}

//GetClearingFeeIndicator reads a ClearingFeeIndicator from NewOrderMultileg.
func (m Message) GetClearingFeeIndicator(f *field.ClearingFeeIndicatorField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//HandlInst is a non-required field for NewOrderMultileg.
func (m Message) HandlInst() (*field.HandlInstField, quickfix.MessageRejectError) {
	f := &field.HandlInstField{}
	err := m.Body.Get(f)
	return f, err
}

//GetHandlInst reads a HandlInst from NewOrderMultileg.
func (m Message) GetHandlInst(f *field.HandlInstField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//ExecInst is a non-required field for NewOrderMultileg.
func (m Message) ExecInst() (*field.ExecInstField, quickfix.MessageRejectError) {
	f := &field.ExecInstField{}
	err := m.Body.Get(f)
	return f, err
}

//GetExecInst reads a ExecInst from NewOrderMultileg.
func (m Message) GetExecInst(f *field.ExecInstField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//MinQty is a non-required field for NewOrderMultileg.
func (m Message) MinQty() (*field.MinQtyField, quickfix.MessageRejectError) {
	f := &field.MinQtyField{}
	err := m.Body.Get(f)
	return f, err
}

//GetMinQty reads a MinQty from NewOrderMultileg.
func (m Message) GetMinQty(f *field.MinQtyField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//MaxFloor is a non-required field for NewOrderMultileg.
func (m Message) MaxFloor() (*field.MaxFloorField, quickfix.MessageRejectError) {
	f := &field.MaxFloorField{}
	err := m.Body.Get(f)
	return f, err
}

//GetMaxFloor reads a MaxFloor from NewOrderMultileg.
func (m Message) GetMaxFloor(f *field.MaxFloorField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//ExDestination is a non-required field for NewOrderMultileg.
func (m Message) ExDestination() (*field.ExDestinationField, quickfix.MessageRejectError) {
	f := &field.ExDestinationField{}
	err := m.Body.Get(f)
	return f, err
}

//GetExDestination reads a ExDestination from NewOrderMultileg.
func (m Message) GetExDestination(f *field.ExDestinationField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//NoTradingSessions is a non-required field for NewOrderMultileg.
func (m Message) NoTradingSessions() (*field.NoTradingSessionsField, quickfix.MessageRejectError) {
	f := &field.NoTradingSessionsField{}
	err := m.Body.Get(f)
	return f, err
}

//GetNoTradingSessions reads a NoTradingSessions from NewOrderMultileg.
func (m Message) GetNoTradingSessions(f *field.NoTradingSessionsField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//ProcessCode is a non-required field for NewOrderMultileg.
func (m Message) ProcessCode() (*field.ProcessCodeField, quickfix.MessageRejectError) {
	f := &field.ProcessCodeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetProcessCode reads a ProcessCode from NewOrderMultileg.
func (m Message) GetProcessCode(f *field.ProcessCodeField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//Side is a required field for NewOrderMultileg.
func (m Message) Side() (*field.SideField, quickfix.MessageRejectError) {
	f := &field.SideField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSide reads a Side from NewOrderMultileg.
func (m Message) GetSide(f *field.SideField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//Symbol is a non-required field for NewOrderMultileg.
func (m Message) Symbol() (*field.SymbolField, quickfix.MessageRejectError) {
	f := &field.SymbolField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSymbol reads a Symbol from NewOrderMultileg.
func (m Message) GetSymbol(f *field.SymbolField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//SymbolSfx is a non-required field for NewOrderMultileg.
func (m Message) SymbolSfx() (*field.SymbolSfxField, quickfix.MessageRejectError) {
	f := &field.SymbolSfxField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSymbolSfx reads a SymbolSfx from NewOrderMultileg.
func (m Message) GetSymbolSfx(f *field.SymbolSfxField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//SecurityID is a non-required field for NewOrderMultileg.
func (m Message) SecurityID() (*field.SecurityIDField, quickfix.MessageRejectError) {
	f := &field.SecurityIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSecurityID reads a SecurityID from NewOrderMultileg.
func (m Message) GetSecurityID(f *field.SecurityIDField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//SecurityIDSource is a non-required field for NewOrderMultileg.
func (m Message) SecurityIDSource() (*field.SecurityIDSourceField, quickfix.MessageRejectError) {
	f := &field.SecurityIDSourceField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSecurityIDSource reads a SecurityIDSource from NewOrderMultileg.
func (m Message) GetSecurityIDSource(f *field.SecurityIDSourceField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//NoSecurityAltID is a non-required field for NewOrderMultileg.
func (m Message) NoSecurityAltID() (*field.NoSecurityAltIDField, quickfix.MessageRejectError) {
	f := &field.NoSecurityAltIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetNoSecurityAltID reads a NoSecurityAltID from NewOrderMultileg.
func (m Message) GetNoSecurityAltID(f *field.NoSecurityAltIDField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//Product is a non-required field for NewOrderMultileg.
func (m Message) Product() (*field.ProductField, quickfix.MessageRejectError) {
	f := &field.ProductField{}
	err := m.Body.Get(f)
	return f, err
}

//GetProduct reads a Product from NewOrderMultileg.
func (m Message) GetProduct(f *field.ProductField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//CFICode is a non-required field for NewOrderMultileg.
func (m Message) CFICode() (*field.CFICodeField, quickfix.MessageRejectError) {
	f := &field.CFICodeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetCFICode reads a CFICode from NewOrderMultileg.
func (m Message) GetCFICode(f *field.CFICodeField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//SecurityType is a non-required field for NewOrderMultileg.
func (m Message) SecurityType() (*field.SecurityTypeField, quickfix.MessageRejectError) {
	f := &field.SecurityTypeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSecurityType reads a SecurityType from NewOrderMultileg.
func (m Message) GetSecurityType(f *field.SecurityTypeField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//SecuritySubType is a non-required field for NewOrderMultileg.
func (m Message) SecuritySubType() (*field.SecuritySubTypeField, quickfix.MessageRejectError) {
	f := &field.SecuritySubTypeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSecuritySubType reads a SecuritySubType from NewOrderMultileg.
func (m Message) GetSecuritySubType(f *field.SecuritySubTypeField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//MaturityMonthYear is a non-required field for NewOrderMultileg.
func (m Message) MaturityMonthYear() (*field.MaturityMonthYearField, quickfix.MessageRejectError) {
	f := &field.MaturityMonthYearField{}
	err := m.Body.Get(f)
	return f, err
}

//GetMaturityMonthYear reads a MaturityMonthYear from NewOrderMultileg.
func (m Message) GetMaturityMonthYear(f *field.MaturityMonthYearField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//MaturityDate is a non-required field for NewOrderMultileg.
func (m Message) MaturityDate() (*field.MaturityDateField, quickfix.MessageRejectError) {
	f := &field.MaturityDateField{}
	err := m.Body.Get(f)
	return f, err
}

//GetMaturityDate reads a MaturityDate from NewOrderMultileg.
func (m Message) GetMaturityDate(f *field.MaturityDateField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//CouponPaymentDate is a non-required field for NewOrderMultileg.
func (m Message) CouponPaymentDate() (*field.CouponPaymentDateField, quickfix.MessageRejectError) {
	f := &field.CouponPaymentDateField{}
	err := m.Body.Get(f)
	return f, err
}

//GetCouponPaymentDate reads a CouponPaymentDate from NewOrderMultileg.
func (m Message) GetCouponPaymentDate(f *field.CouponPaymentDateField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//IssueDate is a non-required field for NewOrderMultileg.
func (m Message) IssueDate() (*field.IssueDateField, quickfix.MessageRejectError) {
	f := &field.IssueDateField{}
	err := m.Body.Get(f)
	return f, err
}

//GetIssueDate reads a IssueDate from NewOrderMultileg.
func (m Message) GetIssueDate(f *field.IssueDateField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//RepoCollateralSecurityType is a non-required field for NewOrderMultileg.
func (m Message) RepoCollateralSecurityType() (*field.RepoCollateralSecurityTypeField, quickfix.MessageRejectError) {
	f := &field.RepoCollateralSecurityTypeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetRepoCollateralSecurityType reads a RepoCollateralSecurityType from NewOrderMultileg.
func (m Message) GetRepoCollateralSecurityType(f *field.RepoCollateralSecurityTypeField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//RepurchaseTerm is a non-required field for NewOrderMultileg.
func (m Message) RepurchaseTerm() (*field.RepurchaseTermField, quickfix.MessageRejectError) {
	f := &field.RepurchaseTermField{}
	err := m.Body.Get(f)
	return f, err
}

//GetRepurchaseTerm reads a RepurchaseTerm from NewOrderMultileg.
func (m Message) GetRepurchaseTerm(f *field.RepurchaseTermField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//RepurchaseRate is a non-required field for NewOrderMultileg.
func (m Message) RepurchaseRate() (*field.RepurchaseRateField, quickfix.MessageRejectError) {
	f := &field.RepurchaseRateField{}
	err := m.Body.Get(f)
	return f, err
}

//GetRepurchaseRate reads a RepurchaseRate from NewOrderMultileg.
func (m Message) GetRepurchaseRate(f *field.RepurchaseRateField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//Factor is a non-required field for NewOrderMultileg.
func (m Message) Factor() (*field.FactorField, quickfix.MessageRejectError) {
	f := &field.FactorField{}
	err := m.Body.Get(f)
	return f, err
}

//GetFactor reads a Factor from NewOrderMultileg.
func (m Message) GetFactor(f *field.FactorField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//CreditRating is a non-required field for NewOrderMultileg.
func (m Message) CreditRating() (*field.CreditRatingField, quickfix.MessageRejectError) {
	f := &field.CreditRatingField{}
	err := m.Body.Get(f)
	return f, err
}

//GetCreditRating reads a CreditRating from NewOrderMultileg.
func (m Message) GetCreditRating(f *field.CreditRatingField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//InstrRegistry is a non-required field for NewOrderMultileg.
func (m Message) InstrRegistry() (*field.InstrRegistryField, quickfix.MessageRejectError) {
	f := &field.InstrRegistryField{}
	err := m.Body.Get(f)
	return f, err
}

//GetInstrRegistry reads a InstrRegistry from NewOrderMultileg.
func (m Message) GetInstrRegistry(f *field.InstrRegistryField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//CountryOfIssue is a non-required field for NewOrderMultileg.
func (m Message) CountryOfIssue() (*field.CountryOfIssueField, quickfix.MessageRejectError) {
	f := &field.CountryOfIssueField{}
	err := m.Body.Get(f)
	return f, err
}

//GetCountryOfIssue reads a CountryOfIssue from NewOrderMultileg.
func (m Message) GetCountryOfIssue(f *field.CountryOfIssueField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//StateOrProvinceOfIssue is a non-required field for NewOrderMultileg.
func (m Message) StateOrProvinceOfIssue() (*field.StateOrProvinceOfIssueField, quickfix.MessageRejectError) {
	f := &field.StateOrProvinceOfIssueField{}
	err := m.Body.Get(f)
	return f, err
}

//GetStateOrProvinceOfIssue reads a StateOrProvinceOfIssue from NewOrderMultileg.
func (m Message) GetStateOrProvinceOfIssue(f *field.StateOrProvinceOfIssueField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//LocaleOfIssue is a non-required field for NewOrderMultileg.
func (m Message) LocaleOfIssue() (*field.LocaleOfIssueField, quickfix.MessageRejectError) {
	f := &field.LocaleOfIssueField{}
	err := m.Body.Get(f)
	return f, err
}

//GetLocaleOfIssue reads a LocaleOfIssue from NewOrderMultileg.
func (m Message) GetLocaleOfIssue(f *field.LocaleOfIssueField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//RedemptionDate is a non-required field for NewOrderMultileg.
func (m Message) RedemptionDate() (*field.RedemptionDateField, quickfix.MessageRejectError) {
	f := &field.RedemptionDateField{}
	err := m.Body.Get(f)
	return f, err
}

//GetRedemptionDate reads a RedemptionDate from NewOrderMultileg.
func (m Message) GetRedemptionDate(f *field.RedemptionDateField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//StrikePrice is a non-required field for NewOrderMultileg.
func (m Message) StrikePrice() (*field.StrikePriceField, quickfix.MessageRejectError) {
	f := &field.StrikePriceField{}
	err := m.Body.Get(f)
	return f, err
}

//GetStrikePrice reads a StrikePrice from NewOrderMultileg.
func (m Message) GetStrikePrice(f *field.StrikePriceField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//StrikeCurrency is a non-required field for NewOrderMultileg.
func (m Message) StrikeCurrency() (*field.StrikeCurrencyField, quickfix.MessageRejectError) {
	f := &field.StrikeCurrencyField{}
	err := m.Body.Get(f)
	return f, err
}

//GetStrikeCurrency reads a StrikeCurrency from NewOrderMultileg.
func (m Message) GetStrikeCurrency(f *field.StrikeCurrencyField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//OptAttribute is a non-required field for NewOrderMultileg.
func (m Message) OptAttribute() (*field.OptAttributeField, quickfix.MessageRejectError) {
	f := &field.OptAttributeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetOptAttribute reads a OptAttribute from NewOrderMultileg.
func (m Message) GetOptAttribute(f *field.OptAttributeField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//ContractMultiplier is a non-required field for NewOrderMultileg.
func (m Message) ContractMultiplier() (*field.ContractMultiplierField, quickfix.MessageRejectError) {
	f := &field.ContractMultiplierField{}
	err := m.Body.Get(f)
	return f, err
}

//GetContractMultiplier reads a ContractMultiplier from NewOrderMultileg.
func (m Message) GetContractMultiplier(f *field.ContractMultiplierField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//CouponRate is a non-required field for NewOrderMultileg.
func (m Message) CouponRate() (*field.CouponRateField, quickfix.MessageRejectError) {
	f := &field.CouponRateField{}
	err := m.Body.Get(f)
	return f, err
}

//GetCouponRate reads a CouponRate from NewOrderMultileg.
func (m Message) GetCouponRate(f *field.CouponRateField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//SecurityExchange is a non-required field for NewOrderMultileg.
func (m Message) SecurityExchange() (*field.SecurityExchangeField, quickfix.MessageRejectError) {
	f := &field.SecurityExchangeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSecurityExchange reads a SecurityExchange from NewOrderMultileg.
func (m Message) GetSecurityExchange(f *field.SecurityExchangeField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//Issuer is a non-required field for NewOrderMultileg.
func (m Message) Issuer() (*field.IssuerField, quickfix.MessageRejectError) {
	f := &field.IssuerField{}
	err := m.Body.Get(f)
	return f, err
}

//GetIssuer reads a Issuer from NewOrderMultileg.
func (m Message) GetIssuer(f *field.IssuerField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//EncodedIssuerLen is a non-required field for NewOrderMultileg.
func (m Message) EncodedIssuerLen() (*field.EncodedIssuerLenField, quickfix.MessageRejectError) {
	f := &field.EncodedIssuerLenField{}
	err := m.Body.Get(f)
	return f, err
}

//GetEncodedIssuerLen reads a EncodedIssuerLen from NewOrderMultileg.
func (m Message) GetEncodedIssuerLen(f *field.EncodedIssuerLenField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//EncodedIssuer is a non-required field for NewOrderMultileg.
func (m Message) EncodedIssuer() (*field.EncodedIssuerField, quickfix.MessageRejectError) {
	f := &field.EncodedIssuerField{}
	err := m.Body.Get(f)
	return f, err
}

//GetEncodedIssuer reads a EncodedIssuer from NewOrderMultileg.
func (m Message) GetEncodedIssuer(f *field.EncodedIssuerField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//SecurityDesc is a non-required field for NewOrderMultileg.
func (m Message) SecurityDesc() (*field.SecurityDescField, quickfix.MessageRejectError) {
	f := &field.SecurityDescField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSecurityDesc reads a SecurityDesc from NewOrderMultileg.
func (m Message) GetSecurityDesc(f *field.SecurityDescField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//EncodedSecurityDescLen is a non-required field for NewOrderMultileg.
func (m Message) EncodedSecurityDescLen() (*field.EncodedSecurityDescLenField, quickfix.MessageRejectError) {
	f := &field.EncodedSecurityDescLenField{}
	err := m.Body.Get(f)
	return f, err
}

//GetEncodedSecurityDescLen reads a EncodedSecurityDescLen from NewOrderMultileg.
func (m Message) GetEncodedSecurityDescLen(f *field.EncodedSecurityDescLenField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//EncodedSecurityDesc is a non-required field for NewOrderMultileg.
func (m Message) EncodedSecurityDesc() (*field.EncodedSecurityDescField, quickfix.MessageRejectError) {
	f := &field.EncodedSecurityDescField{}
	err := m.Body.Get(f)
	return f, err
}

//GetEncodedSecurityDesc reads a EncodedSecurityDesc from NewOrderMultileg.
func (m Message) GetEncodedSecurityDesc(f *field.EncodedSecurityDescField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//Pool is a non-required field for NewOrderMultileg.
func (m Message) Pool() (*field.PoolField, quickfix.MessageRejectError) {
	f := &field.PoolField{}
	err := m.Body.Get(f)
	return f, err
}

//GetPool reads a Pool from NewOrderMultileg.
func (m Message) GetPool(f *field.PoolField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//ContractSettlMonth is a non-required field for NewOrderMultileg.
func (m Message) ContractSettlMonth() (*field.ContractSettlMonthField, quickfix.MessageRejectError) {
	f := &field.ContractSettlMonthField{}
	err := m.Body.Get(f)
	return f, err
}

//GetContractSettlMonth reads a ContractSettlMonth from NewOrderMultileg.
func (m Message) GetContractSettlMonth(f *field.ContractSettlMonthField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//CPProgram is a non-required field for NewOrderMultileg.
func (m Message) CPProgram() (*field.CPProgramField, quickfix.MessageRejectError) {
	f := &field.CPProgramField{}
	err := m.Body.Get(f)
	return f, err
}

//GetCPProgram reads a CPProgram from NewOrderMultileg.
func (m Message) GetCPProgram(f *field.CPProgramField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//CPRegType is a non-required field for NewOrderMultileg.
func (m Message) CPRegType() (*field.CPRegTypeField, quickfix.MessageRejectError) {
	f := &field.CPRegTypeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetCPRegType reads a CPRegType from NewOrderMultileg.
func (m Message) GetCPRegType(f *field.CPRegTypeField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//NoEvents is a non-required field for NewOrderMultileg.
func (m Message) NoEvents() (*field.NoEventsField, quickfix.MessageRejectError) {
	f := &field.NoEventsField{}
	err := m.Body.Get(f)
	return f, err
}

//GetNoEvents reads a NoEvents from NewOrderMultileg.
func (m Message) GetNoEvents(f *field.NoEventsField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//DatedDate is a non-required field for NewOrderMultileg.
func (m Message) DatedDate() (*field.DatedDateField, quickfix.MessageRejectError) {
	f := &field.DatedDateField{}
	err := m.Body.Get(f)
	return f, err
}

//GetDatedDate reads a DatedDate from NewOrderMultileg.
func (m Message) GetDatedDate(f *field.DatedDateField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//InterestAccrualDate is a non-required field for NewOrderMultileg.
func (m Message) InterestAccrualDate() (*field.InterestAccrualDateField, quickfix.MessageRejectError) {
	f := &field.InterestAccrualDateField{}
	err := m.Body.Get(f)
	return f, err
}

//GetInterestAccrualDate reads a InterestAccrualDate from NewOrderMultileg.
func (m Message) GetInterestAccrualDate(f *field.InterestAccrualDateField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//SecurityStatus is a non-required field for NewOrderMultileg.
func (m Message) SecurityStatus() (*field.SecurityStatusField, quickfix.MessageRejectError) {
	f := &field.SecurityStatusField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSecurityStatus reads a SecurityStatus from NewOrderMultileg.
func (m Message) GetSecurityStatus(f *field.SecurityStatusField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//SettleOnOpenFlag is a non-required field for NewOrderMultileg.
func (m Message) SettleOnOpenFlag() (*field.SettleOnOpenFlagField, quickfix.MessageRejectError) {
	f := &field.SettleOnOpenFlagField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSettleOnOpenFlag reads a SettleOnOpenFlag from NewOrderMultileg.
func (m Message) GetSettleOnOpenFlag(f *field.SettleOnOpenFlagField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//InstrmtAssignmentMethod is a non-required field for NewOrderMultileg.
func (m Message) InstrmtAssignmentMethod() (*field.InstrmtAssignmentMethodField, quickfix.MessageRejectError) {
	f := &field.InstrmtAssignmentMethodField{}
	err := m.Body.Get(f)
	return f, err
}

//GetInstrmtAssignmentMethod reads a InstrmtAssignmentMethod from NewOrderMultileg.
func (m Message) GetInstrmtAssignmentMethod(f *field.InstrmtAssignmentMethodField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//StrikeMultiplier is a non-required field for NewOrderMultileg.
func (m Message) StrikeMultiplier() (*field.StrikeMultiplierField, quickfix.MessageRejectError) {
	f := &field.StrikeMultiplierField{}
	err := m.Body.Get(f)
	return f, err
}

//GetStrikeMultiplier reads a StrikeMultiplier from NewOrderMultileg.
func (m Message) GetStrikeMultiplier(f *field.StrikeMultiplierField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//StrikeValue is a non-required field for NewOrderMultileg.
func (m Message) StrikeValue() (*field.StrikeValueField, quickfix.MessageRejectError) {
	f := &field.StrikeValueField{}
	err := m.Body.Get(f)
	return f, err
}

//GetStrikeValue reads a StrikeValue from NewOrderMultileg.
func (m Message) GetStrikeValue(f *field.StrikeValueField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//MinPriceIncrement is a non-required field for NewOrderMultileg.
func (m Message) MinPriceIncrement() (*field.MinPriceIncrementField, quickfix.MessageRejectError) {
	f := &field.MinPriceIncrementField{}
	err := m.Body.Get(f)
	return f, err
}

//GetMinPriceIncrement reads a MinPriceIncrement from NewOrderMultileg.
func (m Message) GetMinPriceIncrement(f *field.MinPriceIncrementField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//PositionLimit is a non-required field for NewOrderMultileg.
func (m Message) PositionLimit() (*field.PositionLimitField, quickfix.MessageRejectError) {
	f := &field.PositionLimitField{}
	err := m.Body.Get(f)
	return f, err
}

//GetPositionLimit reads a PositionLimit from NewOrderMultileg.
func (m Message) GetPositionLimit(f *field.PositionLimitField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//NTPositionLimit is a non-required field for NewOrderMultileg.
func (m Message) NTPositionLimit() (*field.NTPositionLimitField, quickfix.MessageRejectError) {
	f := &field.NTPositionLimitField{}
	err := m.Body.Get(f)
	return f, err
}

//GetNTPositionLimit reads a NTPositionLimit from NewOrderMultileg.
func (m Message) GetNTPositionLimit(f *field.NTPositionLimitField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//NoInstrumentParties is a non-required field for NewOrderMultileg.
func (m Message) NoInstrumentParties() (*field.NoInstrumentPartiesField, quickfix.MessageRejectError) {
	f := &field.NoInstrumentPartiesField{}
	err := m.Body.Get(f)
	return f, err
}

//GetNoInstrumentParties reads a NoInstrumentParties from NewOrderMultileg.
func (m Message) GetNoInstrumentParties(f *field.NoInstrumentPartiesField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//UnitOfMeasure is a non-required field for NewOrderMultileg.
func (m Message) UnitOfMeasure() (*field.UnitOfMeasureField, quickfix.MessageRejectError) {
	f := &field.UnitOfMeasureField{}
	err := m.Body.Get(f)
	return f, err
}

//GetUnitOfMeasure reads a UnitOfMeasure from NewOrderMultileg.
func (m Message) GetUnitOfMeasure(f *field.UnitOfMeasureField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//TimeUnit is a non-required field for NewOrderMultileg.
func (m Message) TimeUnit() (*field.TimeUnitField, quickfix.MessageRejectError) {
	f := &field.TimeUnitField{}
	err := m.Body.Get(f)
	return f, err
}

//GetTimeUnit reads a TimeUnit from NewOrderMultileg.
func (m Message) GetTimeUnit(f *field.TimeUnitField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//MaturityTime is a non-required field for NewOrderMultileg.
func (m Message) MaturityTime() (*field.MaturityTimeField, quickfix.MessageRejectError) {
	f := &field.MaturityTimeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetMaturityTime reads a MaturityTime from NewOrderMultileg.
func (m Message) GetMaturityTime(f *field.MaturityTimeField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//SecurityGroup is a non-required field for NewOrderMultileg.
func (m Message) SecurityGroup() (*field.SecurityGroupField, quickfix.MessageRejectError) {
	f := &field.SecurityGroupField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSecurityGroup reads a SecurityGroup from NewOrderMultileg.
func (m Message) GetSecurityGroup(f *field.SecurityGroupField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//MinPriceIncrementAmount is a non-required field for NewOrderMultileg.
func (m Message) MinPriceIncrementAmount() (*field.MinPriceIncrementAmountField, quickfix.MessageRejectError) {
	f := &field.MinPriceIncrementAmountField{}
	err := m.Body.Get(f)
	return f, err
}

//GetMinPriceIncrementAmount reads a MinPriceIncrementAmount from NewOrderMultileg.
func (m Message) GetMinPriceIncrementAmount(f *field.MinPriceIncrementAmountField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//UnitOfMeasureQty is a non-required field for NewOrderMultileg.
func (m Message) UnitOfMeasureQty() (*field.UnitOfMeasureQtyField, quickfix.MessageRejectError) {
	f := &field.UnitOfMeasureQtyField{}
	err := m.Body.Get(f)
	return f, err
}

//GetUnitOfMeasureQty reads a UnitOfMeasureQty from NewOrderMultileg.
func (m Message) GetUnitOfMeasureQty(f *field.UnitOfMeasureQtyField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//SecurityXMLLen is a non-required field for NewOrderMultileg.
func (m Message) SecurityXMLLen() (*field.SecurityXMLLenField, quickfix.MessageRejectError) {
	f := &field.SecurityXMLLenField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSecurityXMLLen reads a SecurityXMLLen from NewOrderMultileg.
func (m Message) GetSecurityXMLLen(f *field.SecurityXMLLenField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//SecurityXML is a non-required field for NewOrderMultileg.
func (m Message) SecurityXML() (*field.SecurityXMLField, quickfix.MessageRejectError) {
	f := &field.SecurityXMLField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSecurityXML reads a SecurityXML from NewOrderMultileg.
func (m Message) GetSecurityXML(f *field.SecurityXMLField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//SecurityXMLSchema is a non-required field for NewOrderMultileg.
func (m Message) SecurityXMLSchema() (*field.SecurityXMLSchemaField, quickfix.MessageRejectError) {
	f := &field.SecurityXMLSchemaField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSecurityXMLSchema reads a SecurityXMLSchema from NewOrderMultileg.
func (m Message) GetSecurityXMLSchema(f *field.SecurityXMLSchemaField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//ProductComplex is a non-required field for NewOrderMultileg.
func (m Message) ProductComplex() (*field.ProductComplexField, quickfix.MessageRejectError) {
	f := &field.ProductComplexField{}
	err := m.Body.Get(f)
	return f, err
}

//GetProductComplex reads a ProductComplex from NewOrderMultileg.
func (m Message) GetProductComplex(f *field.ProductComplexField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//PriceUnitOfMeasure is a non-required field for NewOrderMultileg.
func (m Message) PriceUnitOfMeasure() (*field.PriceUnitOfMeasureField, quickfix.MessageRejectError) {
	f := &field.PriceUnitOfMeasureField{}
	err := m.Body.Get(f)
	return f, err
}

//GetPriceUnitOfMeasure reads a PriceUnitOfMeasure from NewOrderMultileg.
func (m Message) GetPriceUnitOfMeasure(f *field.PriceUnitOfMeasureField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//PriceUnitOfMeasureQty is a non-required field for NewOrderMultileg.
func (m Message) PriceUnitOfMeasureQty() (*field.PriceUnitOfMeasureQtyField, quickfix.MessageRejectError) {
	f := &field.PriceUnitOfMeasureQtyField{}
	err := m.Body.Get(f)
	return f, err
}

//GetPriceUnitOfMeasureQty reads a PriceUnitOfMeasureQty from NewOrderMultileg.
func (m Message) GetPriceUnitOfMeasureQty(f *field.PriceUnitOfMeasureQtyField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//SettlMethod is a non-required field for NewOrderMultileg.
func (m Message) SettlMethod() (*field.SettlMethodField, quickfix.MessageRejectError) {
	f := &field.SettlMethodField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSettlMethod reads a SettlMethod from NewOrderMultileg.
func (m Message) GetSettlMethod(f *field.SettlMethodField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//ExerciseStyle is a non-required field for NewOrderMultileg.
func (m Message) ExerciseStyle() (*field.ExerciseStyleField, quickfix.MessageRejectError) {
	f := &field.ExerciseStyleField{}
	err := m.Body.Get(f)
	return f, err
}

//GetExerciseStyle reads a ExerciseStyle from NewOrderMultileg.
func (m Message) GetExerciseStyle(f *field.ExerciseStyleField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//OptPayoutAmount is a non-required field for NewOrderMultileg.
func (m Message) OptPayoutAmount() (*field.OptPayoutAmountField, quickfix.MessageRejectError) {
	f := &field.OptPayoutAmountField{}
	err := m.Body.Get(f)
	return f, err
}

//GetOptPayoutAmount reads a OptPayoutAmount from NewOrderMultileg.
func (m Message) GetOptPayoutAmount(f *field.OptPayoutAmountField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//PriceQuoteMethod is a non-required field for NewOrderMultileg.
func (m Message) PriceQuoteMethod() (*field.PriceQuoteMethodField, quickfix.MessageRejectError) {
	f := &field.PriceQuoteMethodField{}
	err := m.Body.Get(f)
	return f, err
}

//GetPriceQuoteMethod reads a PriceQuoteMethod from NewOrderMultileg.
func (m Message) GetPriceQuoteMethod(f *field.PriceQuoteMethodField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//ListMethod is a non-required field for NewOrderMultileg.
func (m Message) ListMethod() (*field.ListMethodField, quickfix.MessageRejectError) {
	f := &field.ListMethodField{}
	err := m.Body.Get(f)
	return f, err
}

//GetListMethod reads a ListMethod from NewOrderMultileg.
func (m Message) GetListMethod(f *field.ListMethodField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//CapPrice is a non-required field for NewOrderMultileg.
func (m Message) CapPrice() (*field.CapPriceField, quickfix.MessageRejectError) {
	f := &field.CapPriceField{}
	err := m.Body.Get(f)
	return f, err
}

//GetCapPrice reads a CapPrice from NewOrderMultileg.
func (m Message) GetCapPrice(f *field.CapPriceField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//FloorPrice is a non-required field for NewOrderMultileg.
func (m Message) FloorPrice() (*field.FloorPriceField, quickfix.MessageRejectError) {
	f := &field.FloorPriceField{}
	err := m.Body.Get(f)
	return f, err
}

//GetFloorPrice reads a FloorPrice from NewOrderMultileg.
func (m Message) GetFloorPrice(f *field.FloorPriceField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//PutOrCall is a non-required field for NewOrderMultileg.
func (m Message) PutOrCall() (*field.PutOrCallField, quickfix.MessageRejectError) {
	f := &field.PutOrCallField{}
	err := m.Body.Get(f)
	return f, err
}

//GetPutOrCall reads a PutOrCall from NewOrderMultileg.
func (m Message) GetPutOrCall(f *field.PutOrCallField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//FlexibleIndicator is a non-required field for NewOrderMultileg.
func (m Message) FlexibleIndicator() (*field.FlexibleIndicatorField, quickfix.MessageRejectError) {
	f := &field.FlexibleIndicatorField{}
	err := m.Body.Get(f)
	return f, err
}

//GetFlexibleIndicator reads a FlexibleIndicator from NewOrderMultileg.
func (m Message) GetFlexibleIndicator(f *field.FlexibleIndicatorField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//FlexProductEligibilityIndicator is a non-required field for NewOrderMultileg.
func (m Message) FlexProductEligibilityIndicator() (*field.FlexProductEligibilityIndicatorField, quickfix.MessageRejectError) {
	f := &field.FlexProductEligibilityIndicatorField{}
	err := m.Body.Get(f)
	return f, err
}

//GetFlexProductEligibilityIndicator reads a FlexProductEligibilityIndicator from NewOrderMultileg.
func (m Message) GetFlexProductEligibilityIndicator(f *field.FlexProductEligibilityIndicatorField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//ValuationMethod is a non-required field for NewOrderMultileg.
func (m Message) ValuationMethod() (*field.ValuationMethodField, quickfix.MessageRejectError) {
	f := &field.ValuationMethodField{}
	err := m.Body.Get(f)
	return f, err
}

//GetValuationMethod reads a ValuationMethod from NewOrderMultileg.
func (m Message) GetValuationMethod(f *field.ValuationMethodField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//ContractMultiplierUnit is a non-required field for NewOrderMultileg.
func (m Message) ContractMultiplierUnit() (*field.ContractMultiplierUnitField, quickfix.MessageRejectError) {
	f := &field.ContractMultiplierUnitField{}
	err := m.Body.Get(f)
	return f, err
}

//GetContractMultiplierUnit reads a ContractMultiplierUnit from NewOrderMultileg.
func (m Message) GetContractMultiplierUnit(f *field.ContractMultiplierUnitField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//FlowScheduleType is a non-required field for NewOrderMultileg.
func (m Message) FlowScheduleType() (*field.FlowScheduleTypeField, quickfix.MessageRejectError) {
	f := &field.FlowScheduleTypeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetFlowScheduleType reads a FlowScheduleType from NewOrderMultileg.
func (m Message) GetFlowScheduleType(f *field.FlowScheduleTypeField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//RestructuringType is a non-required field for NewOrderMultileg.
func (m Message) RestructuringType() (*field.RestructuringTypeField, quickfix.MessageRejectError) {
	f := &field.RestructuringTypeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetRestructuringType reads a RestructuringType from NewOrderMultileg.
func (m Message) GetRestructuringType(f *field.RestructuringTypeField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//Seniority is a non-required field for NewOrderMultileg.
func (m Message) Seniority() (*field.SeniorityField, quickfix.MessageRejectError) {
	f := &field.SeniorityField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSeniority reads a Seniority from NewOrderMultileg.
func (m Message) GetSeniority(f *field.SeniorityField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//NotionalPercentageOutstanding is a non-required field for NewOrderMultileg.
func (m Message) NotionalPercentageOutstanding() (*field.NotionalPercentageOutstandingField, quickfix.MessageRejectError) {
	f := &field.NotionalPercentageOutstandingField{}
	err := m.Body.Get(f)
	return f, err
}

//GetNotionalPercentageOutstanding reads a NotionalPercentageOutstanding from NewOrderMultileg.
func (m Message) GetNotionalPercentageOutstanding(f *field.NotionalPercentageOutstandingField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//OriginalNotionalPercentageOutstanding is a non-required field for NewOrderMultileg.
func (m Message) OriginalNotionalPercentageOutstanding() (*field.OriginalNotionalPercentageOutstandingField, quickfix.MessageRejectError) {
	f := &field.OriginalNotionalPercentageOutstandingField{}
	err := m.Body.Get(f)
	return f, err
}

//GetOriginalNotionalPercentageOutstanding reads a OriginalNotionalPercentageOutstanding from NewOrderMultileg.
func (m Message) GetOriginalNotionalPercentageOutstanding(f *field.OriginalNotionalPercentageOutstandingField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//AttachmentPoint is a non-required field for NewOrderMultileg.
func (m Message) AttachmentPoint() (*field.AttachmentPointField, quickfix.MessageRejectError) {
	f := &field.AttachmentPointField{}
	err := m.Body.Get(f)
	return f, err
}

//GetAttachmentPoint reads a AttachmentPoint from NewOrderMultileg.
func (m Message) GetAttachmentPoint(f *field.AttachmentPointField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//DetachmentPoint is a non-required field for NewOrderMultileg.
func (m Message) DetachmentPoint() (*field.DetachmentPointField, quickfix.MessageRejectError) {
	f := &field.DetachmentPointField{}
	err := m.Body.Get(f)
	return f, err
}

//GetDetachmentPoint reads a DetachmentPoint from NewOrderMultileg.
func (m Message) GetDetachmentPoint(f *field.DetachmentPointField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//StrikePriceDeterminationMethod is a non-required field for NewOrderMultileg.
func (m Message) StrikePriceDeterminationMethod() (*field.StrikePriceDeterminationMethodField, quickfix.MessageRejectError) {
	f := &field.StrikePriceDeterminationMethodField{}
	err := m.Body.Get(f)
	return f, err
}

//GetStrikePriceDeterminationMethod reads a StrikePriceDeterminationMethod from NewOrderMultileg.
func (m Message) GetStrikePriceDeterminationMethod(f *field.StrikePriceDeterminationMethodField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//StrikePriceBoundaryMethod is a non-required field for NewOrderMultileg.
func (m Message) StrikePriceBoundaryMethod() (*field.StrikePriceBoundaryMethodField, quickfix.MessageRejectError) {
	f := &field.StrikePriceBoundaryMethodField{}
	err := m.Body.Get(f)
	return f, err
}

//GetStrikePriceBoundaryMethod reads a StrikePriceBoundaryMethod from NewOrderMultileg.
func (m Message) GetStrikePriceBoundaryMethod(f *field.StrikePriceBoundaryMethodField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//StrikePriceBoundaryPrecision is a non-required field for NewOrderMultileg.
func (m Message) StrikePriceBoundaryPrecision() (*field.StrikePriceBoundaryPrecisionField, quickfix.MessageRejectError) {
	f := &field.StrikePriceBoundaryPrecisionField{}
	err := m.Body.Get(f)
	return f, err
}

//GetStrikePriceBoundaryPrecision reads a StrikePriceBoundaryPrecision from NewOrderMultileg.
func (m Message) GetStrikePriceBoundaryPrecision(f *field.StrikePriceBoundaryPrecisionField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//UnderlyingPriceDeterminationMethod is a non-required field for NewOrderMultileg.
func (m Message) UnderlyingPriceDeterminationMethod() (*field.UnderlyingPriceDeterminationMethodField, quickfix.MessageRejectError) {
	f := &field.UnderlyingPriceDeterminationMethodField{}
	err := m.Body.Get(f)
	return f, err
}

//GetUnderlyingPriceDeterminationMethod reads a UnderlyingPriceDeterminationMethod from NewOrderMultileg.
func (m Message) GetUnderlyingPriceDeterminationMethod(f *field.UnderlyingPriceDeterminationMethodField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//OptPayoutType is a non-required field for NewOrderMultileg.
func (m Message) OptPayoutType() (*field.OptPayoutTypeField, quickfix.MessageRejectError) {
	f := &field.OptPayoutTypeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetOptPayoutType reads a OptPayoutType from NewOrderMultileg.
func (m Message) GetOptPayoutType(f *field.OptPayoutTypeField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//NoComplexEvents is a non-required field for NewOrderMultileg.
func (m Message) NoComplexEvents() (*field.NoComplexEventsField, quickfix.MessageRejectError) {
	f := &field.NoComplexEventsField{}
	err := m.Body.Get(f)
	return f, err
}

//GetNoComplexEvents reads a NoComplexEvents from NewOrderMultileg.
func (m Message) GetNoComplexEvents(f *field.NoComplexEventsField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//NoUnderlyings is a non-required field for NewOrderMultileg.
func (m Message) NoUnderlyings() (*field.NoUnderlyingsField, quickfix.MessageRejectError) {
	f := &field.NoUnderlyingsField{}
	err := m.Body.Get(f)
	return f, err
}

//GetNoUnderlyings reads a NoUnderlyings from NewOrderMultileg.
func (m Message) GetNoUnderlyings(f *field.NoUnderlyingsField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//PrevClosePx is a non-required field for NewOrderMultileg.
func (m Message) PrevClosePx() (*field.PrevClosePxField, quickfix.MessageRejectError) {
	f := &field.PrevClosePxField{}
	err := m.Body.Get(f)
	return f, err
}

//GetPrevClosePx reads a PrevClosePx from NewOrderMultileg.
func (m Message) GetPrevClosePx(f *field.PrevClosePxField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//NoLegs is a required field for NewOrderMultileg.
func (m Message) NoLegs() (*field.NoLegsField, quickfix.MessageRejectError) {
	f := &field.NoLegsField{}
	err := m.Body.Get(f)
	return f, err
}

//GetNoLegs reads a NoLegs from NewOrderMultileg.
func (m Message) GetNoLegs(f *field.NoLegsField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//LocateReqd is a non-required field for NewOrderMultileg.
func (m Message) LocateReqd() (*field.LocateReqdField, quickfix.MessageRejectError) {
	f := &field.LocateReqdField{}
	err := m.Body.Get(f)
	return f, err
}

//GetLocateReqd reads a LocateReqd from NewOrderMultileg.
func (m Message) GetLocateReqd(f *field.LocateReqdField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//TransactTime is a required field for NewOrderMultileg.
func (m Message) TransactTime() (*field.TransactTimeField, quickfix.MessageRejectError) {
	f := &field.TransactTimeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetTransactTime reads a TransactTime from NewOrderMultileg.
func (m Message) GetTransactTime(f *field.TransactTimeField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//QtyType is a non-required field for NewOrderMultileg.
func (m Message) QtyType() (*field.QtyTypeField, quickfix.MessageRejectError) {
	f := &field.QtyTypeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetQtyType reads a QtyType from NewOrderMultileg.
func (m Message) GetQtyType(f *field.QtyTypeField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//OrderQty is a non-required field for NewOrderMultileg.
func (m Message) OrderQty() (*field.OrderQtyField, quickfix.MessageRejectError) {
	f := &field.OrderQtyField{}
	err := m.Body.Get(f)
	return f, err
}

//GetOrderQty reads a OrderQty from NewOrderMultileg.
func (m Message) GetOrderQty(f *field.OrderQtyField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//CashOrderQty is a non-required field for NewOrderMultileg.
func (m Message) CashOrderQty() (*field.CashOrderQtyField, quickfix.MessageRejectError) {
	f := &field.CashOrderQtyField{}
	err := m.Body.Get(f)
	return f, err
}

//GetCashOrderQty reads a CashOrderQty from NewOrderMultileg.
func (m Message) GetCashOrderQty(f *field.CashOrderQtyField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//OrderPercent is a non-required field for NewOrderMultileg.
func (m Message) OrderPercent() (*field.OrderPercentField, quickfix.MessageRejectError) {
	f := &field.OrderPercentField{}
	err := m.Body.Get(f)
	return f, err
}

//GetOrderPercent reads a OrderPercent from NewOrderMultileg.
func (m Message) GetOrderPercent(f *field.OrderPercentField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//RoundingDirection is a non-required field for NewOrderMultileg.
func (m Message) RoundingDirection() (*field.RoundingDirectionField, quickfix.MessageRejectError) {
	f := &field.RoundingDirectionField{}
	err := m.Body.Get(f)
	return f, err
}

//GetRoundingDirection reads a RoundingDirection from NewOrderMultileg.
func (m Message) GetRoundingDirection(f *field.RoundingDirectionField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//RoundingModulus is a non-required field for NewOrderMultileg.
func (m Message) RoundingModulus() (*field.RoundingModulusField, quickfix.MessageRejectError) {
	f := &field.RoundingModulusField{}
	err := m.Body.Get(f)
	return f, err
}

//GetRoundingModulus reads a RoundingModulus from NewOrderMultileg.
func (m Message) GetRoundingModulus(f *field.RoundingModulusField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//OrdType is a required field for NewOrderMultileg.
func (m Message) OrdType() (*field.OrdTypeField, quickfix.MessageRejectError) {
	f := &field.OrdTypeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetOrdType reads a OrdType from NewOrderMultileg.
func (m Message) GetOrdType(f *field.OrdTypeField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//PriceType is a non-required field for NewOrderMultileg.
func (m Message) PriceType() (*field.PriceTypeField, quickfix.MessageRejectError) {
	f := &field.PriceTypeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetPriceType reads a PriceType from NewOrderMultileg.
func (m Message) GetPriceType(f *field.PriceTypeField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//Price is a non-required field for NewOrderMultileg.
func (m Message) Price() (*field.PriceField, quickfix.MessageRejectError) {
	f := &field.PriceField{}
	err := m.Body.Get(f)
	return f, err
}

//GetPrice reads a Price from NewOrderMultileg.
func (m Message) GetPrice(f *field.PriceField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//StopPx is a non-required field for NewOrderMultileg.
func (m Message) StopPx() (*field.StopPxField, quickfix.MessageRejectError) {
	f := &field.StopPxField{}
	err := m.Body.Get(f)
	return f, err
}

//GetStopPx reads a StopPx from NewOrderMultileg.
func (m Message) GetStopPx(f *field.StopPxField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//Currency is a non-required field for NewOrderMultileg.
func (m Message) Currency() (*field.CurrencyField, quickfix.MessageRejectError) {
	f := &field.CurrencyField{}
	err := m.Body.Get(f)
	return f, err
}

//GetCurrency reads a Currency from NewOrderMultileg.
func (m Message) GetCurrency(f *field.CurrencyField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//ComplianceID is a non-required field for NewOrderMultileg.
func (m Message) ComplianceID() (*field.ComplianceIDField, quickfix.MessageRejectError) {
	f := &field.ComplianceIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetComplianceID reads a ComplianceID from NewOrderMultileg.
func (m Message) GetComplianceID(f *field.ComplianceIDField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//SolicitedFlag is a non-required field for NewOrderMultileg.
func (m Message) SolicitedFlag() (*field.SolicitedFlagField, quickfix.MessageRejectError) {
	f := &field.SolicitedFlagField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSolicitedFlag reads a SolicitedFlag from NewOrderMultileg.
func (m Message) GetSolicitedFlag(f *field.SolicitedFlagField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//IOIID is a non-required field for NewOrderMultileg.
func (m Message) IOIID() (*field.IOIIDField, quickfix.MessageRejectError) {
	f := &field.IOIIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetIOIID reads a IOIID from NewOrderMultileg.
func (m Message) GetIOIID(f *field.IOIIDField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//QuoteID is a non-required field for NewOrderMultileg.
func (m Message) QuoteID() (*field.QuoteIDField, quickfix.MessageRejectError) {
	f := &field.QuoteIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetQuoteID reads a QuoteID from NewOrderMultileg.
func (m Message) GetQuoteID(f *field.QuoteIDField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//TimeInForce is a non-required field for NewOrderMultileg.
func (m Message) TimeInForce() (*field.TimeInForceField, quickfix.MessageRejectError) {
	f := &field.TimeInForceField{}
	err := m.Body.Get(f)
	return f, err
}

//GetTimeInForce reads a TimeInForce from NewOrderMultileg.
func (m Message) GetTimeInForce(f *field.TimeInForceField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//EffectiveTime is a non-required field for NewOrderMultileg.
func (m Message) EffectiveTime() (*field.EffectiveTimeField, quickfix.MessageRejectError) {
	f := &field.EffectiveTimeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetEffectiveTime reads a EffectiveTime from NewOrderMultileg.
func (m Message) GetEffectiveTime(f *field.EffectiveTimeField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//ExpireDate is a non-required field for NewOrderMultileg.
func (m Message) ExpireDate() (*field.ExpireDateField, quickfix.MessageRejectError) {
	f := &field.ExpireDateField{}
	err := m.Body.Get(f)
	return f, err
}

//GetExpireDate reads a ExpireDate from NewOrderMultileg.
func (m Message) GetExpireDate(f *field.ExpireDateField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//ExpireTime is a non-required field for NewOrderMultileg.
func (m Message) ExpireTime() (*field.ExpireTimeField, quickfix.MessageRejectError) {
	f := &field.ExpireTimeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetExpireTime reads a ExpireTime from NewOrderMultileg.
func (m Message) GetExpireTime(f *field.ExpireTimeField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//GTBookingInst is a non-required field for NewOrderMultileg.
func (m Message) GTBookingInst() (*field.GTBookingInstField, quickfix.MessageRejectError) {
	f := &field.GTBookingInstField{}
	err := m.Body.Get(f)
	return f, err
}

//GetGTBookingInst reads a GTBookingInst from NewOrderMultileg.
func (m Message) GetGTBookingInst(f *field.GTBookingInstField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//Commission is a non-required field for NewOrderMultileg.
func (m Message) Commission() (*field.CommissionField, quickfix.MessageRejectError) {
	f := &field.CommissionField{}
	err := m.Body.Get(f)
	return f, err
}

//GetCommission reads a Commission from NewOrderMultileg.
func (m Message) GetCommission(f *field.CommissionField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//CommType is a non-required field for NewOrderMultileg.
func (m Message) CommType() (*field.CommTypeField, quickfix.MessageRejectError) {
	f := &field.CommTypeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetCommType reads a CommType from NewOrderMultileg.
func (m Message) GetCommType(f *field.CommTypeField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//CommCurrency is a non-required field for NewOrderMultileg.
func (m Message) CommCurrency() (*field.CommCurrencyField, quickfix.MessageRejectError) {
	f := &field.CommCurrencyField{}
	err := m.Body.Get(f)
	return f, err
}

//GetCommCurrency reads a CommCurrency from NewOrderMultileg.
func (m Message) GetCommCurrency(f *field.CommCurrencyField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//FundRenewWaiv is a non-required field for NewOrderMultileg.
func (m Message) FundRenewWaiv() (*field.FundRenewWaivField, quickfix.MessageRejectError) {
	f := &field.FundRenewWaivField{}
	err := m.Body.Get(f)
	return f, err
}

//GetFundRenewWaiv reads a FundRenewWaiv from NewOrderMultileg.
func (m Message) GetFundRenewWaiv(f *field.FundRenewWaivField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//OrderCapacity is a non-required field for NewOrderMultileg.
func (m Message) OrderCapacity() (*field.OrderCapacityField, quickfix.MessageRejectError) {
	f := &field.OrderCapacityField{}
	err := m.Body.Get(f)
	return f, err
}

//GetOrderCapacity reads a OrderCapacity from NewOrderMultileg.
func (m Message) GetOrderCapacity(f *field.OrderCapacityField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//OrderRestrictions is a non-required field for NewOrderMultileg.
func (m Message) OrderRestrictions() (*field.OrderRestrictionsField, quickfix.MessageRejectError) {
	f := &field.OrderRestrictionsField{}
	err := m.Body.Get(f)
	return f, err
}

//GetOrderRestrictions reads a OrderRestrictions from NewOrderMultileg.
func (m Message) GetOrderRestrictions(f *field.OrderRestrictionsField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//CustOrderCapacity is a non-required field for NewOrderMultileg.
func (m Message) CustOrderCapacity() (*field.CustOrderCapacityField, quickfix.MessageRejectError) {
	f := &field.CustOrderCapacityField{}
	err := m.Body.Get(f)
	return f, err
}

//GetCustOrderCapacity reads a CustOrderCapacity from NewOrderMultileg.
func (m Message) GetCustOrderCapacity(f *field.CustOrderCapacityField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//ForexReq is a non-required field for NewOrderMultileg.
func (m Message) ForexReq() (*field.ForexReqField, quickfix.MessageRejectError) {
	f := &field.ForexReqField{}
	err := m.Body.Get(f)
	return f, err
}

//GetForexReq reads a ForexReq from NewOrderMultileg.
func (m Message) GetForexReq(f *field.ForexReqField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//SettlCurrency is a non-required field for NewOrderMultileg.
func (m Message) SettlCurrency() (*field.SettlCurrencyField, quickfix.MessageRejectError) {
	f := &field.SettlCurrencyField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSettlCurrency reads a SettlCurrency from NewOrderMultileg.
func (m Message) GetSettlCurrency(f *field.SettlCurrencyField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//BookingType is a non-required field for NewOrderMultileg.
func (m Message) BookingType() (*field.BookingTypeField, quickfix.MessageRejectError) {
	f := &field.BookingTypeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetBookingType reads a BookingType from NewOrderMultileg.
func (m Message) GetBookingType(f *field.BookingTypeField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//Text is a non-required field for NewOrderMultileg.
func (m Message) Text() (*field.TextField, quickfix.MessageRejectError) {
	f := &field.TextField{}
	err := m.Body.Get(f)
	return f, err
}

//GetText reads a Text from NewOrderMultileg.
func (m Message) GetText(f *field.TextField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//EncodedTextLen is a non-required field for NewOrderMultileg.
func (m Message) EncodedTextLen() (*field.EncodedTextLenField, quickfix.MessageRejectError) {
	f := &field.EncodedTextLenField{}
	err := m.Body.Get(f)
	return f, err
}

//GetEncodedTextLen reads a EncodedTextLen from NewOrderMultileg.
func (m Message) GetEncodedTextLen(f *field.EncodedTextLenField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//EncodedText is a non-required field for NewOrderMultileg.
func (m Message) EncodedText() (*field.EncodedTextField, quickfix.MessageRejectError) {
	f := &field.EncodedTextField{}
	err := m.Body.Get(f)
	return f, err
}

//GetEncodedText reads a EncodedText from NewOrderMultileg.
func (m Message) GetEncodedText(f *field.EncodedTextField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//PositionEffect is a non-required field for NewOrderMultileg.
func (m Message) PositionEffect() (*field.PositionEffectField, quickfix.MessageRejectError) {
	f := &field.PositionEffectField{}
	err := m.Body.Get(f)
	return f, err
}

//GetPositionEffect reads a PositionEffect from NewOrderMultileg.
func (m Message) GetPositionEffect(f *field.PositionEffectField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//CoveredOrUncovered is a non-required field for NewOrderMultileg.
func (m Message) CoveredOrUncovered() (*field.CoveredOrUncoveredField, quickfix.MessageRejectError) {
	f := &field.CoveredOrUncoveredField{}
	err := m.Body.Get(f)
	return f, err
}

//GetCoveredOrUncovered reads a CoveredOrUncovered from NewOrderMultileg.
func (m Message) GetCoveredOrUncovered(f *field.CoveredOrUncoveredField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//MaxShow is a non-required field for NewOrderMultileg.
func (m Message) MaxShow() (*field.MaxShowField, quickfix.MessageRejectError) {
	f := &field.MaxShowField{}
	err := m.Body.Get(f)
	return f, err
}

//GetMaxShow reads a MaxShow from NewOrderMultileg.
func (m Message) GetMaxShow(f *field.MaxShowField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//PegOffsetValue is a non-required field for NewOrderMultileg.
func (m Message) PegOffsetValue() (*field.PegOffsetValueField, quickfix.MessageRejectError) {
	f := &field.PegOffsetValueField{}
	err := m.Body.Get(f)
	return f, err
}

//GetPegOffsetValue reads a PegOffsetValue from NewOrderMultileg.
func (m Message) GetPegOffsetValue(f *field.PegOffsetValueField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//PegMoveType is a non-required field for NewOrderMultileg.
func (m Message) PegMoveType() (*field.PegMoveTypeField, quickfix.MessageRejectError) {
	f := &field.PegMoveTypeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetPegMoveType reads a PegMoveType from NewOrderMultileg.
func (m Message) GetPegMoveType(f *field.PegMoveTypeField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//PegOffsetType is a non-required field for NewOrderMultileg.
func (m Message) PegOffsetType() (*field.PegOffsetTypeField, quickfix.MessageRejectError) {
	f := &field.PegOffsetTypeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetPegOffsetType reads a PegOffsetType from NewOrderMultileg.
func (m Message) GetPegOffsetType(f *field.PegOffsetTypeField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//PegLimitType is a non-required field for NewOrderMultileg.
func (m Message) PegLimitType() (*field.PegLimitTypeField, quickfix.MessageRejectError) {
	f := &field.PegLimitTypeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetPegLimitType reads a PegLimitType from NewOrderMultileg.
func (m Message) GetPegLimitType(f *field.PegLimitTypeField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//PegRoundDirection is a non-required field for NewOrderMultileg.
func (m Message) PegRoundDirection() (*field.PegRoundDirectionField, quickfix.MessageRejectError) {
	f := &field.PegRoundDirectionField{}
	err := m.Body.Get(f)
	return f, err
}

//GetPegRoundDirection reads a PegRoundDirection from NewOrderMultileg.
func (m Message) GetPegRoundDirection(f *field.PegRoundDirectionField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//PegScope is a non-required field for NewOrderMultileg.
func (m Message) PegScope() (*field.PegScopeField, quickfix.MessageRejectError) {
	f := &field.PegScopeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetPegScope reads a PegScope from NewOrderMultileg.
func (m Message) GetPegScope(f *field.PegScopeField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//PegPriceType is a non-required field for NewOrderMultileg.
func (m Message) PegPriceType() (*field.PegPriceTypeField, quickfix.MessageRejectError) {
	f := &field.PegPriceTypeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetPegPriceType reads a PegPriceType from NewOrderMultileg.
func (m Message) GetPegPriceType(f *field.PegPriceTypeField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//PegSecurityIDSource is a non-required field for NewOrderMultileg.
func (m Message) PegSecurityIDSource() (*field.PegSecurityIDSourceField, quickfix.MessageRejectError) {
	f := &field.PegSecurityIDSourceField{}
	err := m.Body.Get(f)
	return f, err
}

//GetPegSecurityIDSource reads a PegSecurityIDSource from NewOrderMultileg.
func (m Message) GetPegSecurityIDSource(f *field.PegSecurityIDSourceField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//PegSecurityID is a non-required field for NewOrderMultileg.
func (m Message) PegSecurityID() (*field.PegSecurityIDField, quickfix.MessageRejectError) {
	f := &field.PegSecurityIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetPegSecurityID reads a PegSecurityID from NewOrderMultileg.
func (m Message) GetPegSecurityID(f *field.PegSecurityIDField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//PegSymbol is a non-required field for NewOrderMultileg.
func (m Message) PegSymbol() (*field.PegSymbolField, quickfix.MessageRejectError) {
	f := &field.PegSymbolField{}
	err := m.Body.Get(f)
	return f, err
}

//GetPegSymbol reads a PegSymbol from NewOrderMultileg.
func (m Message) GetPegSymbol(f *field.PegSymbolField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//PegSecurityDesc is a non-required field for NewOrderMultileg.
func (m Message) PegSecurityDesc() (*field.PegSecurityDescField, quickfix.MessageRejectError) {
	f := &field.PegSecurityDescField{}
	err := m.Body.Get(f)
	return f, err
}

//GetPegSecurityDesc reads a PegSecurityDesc from NewOrderMultileg.
func (m Message) GetPegSecurityDesc(f *field.PegSecurityDescField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//DiscretionInst is a non-required field for NewOrderMultileg.
func (m Message) DiscretionInst() (*field.DiscretionInstField, quickfix.MessageRejectError) {
	f := &field.DiscretionInstField{}
	err := m.Body.Get(f)
	return f, err
}

//GetDiscretionInst reads a DiscretionInst from NewOrderMultileg.
func (m Message) GetDiscretionInst(f *field.DiscretionInstField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//DiscretionOffsetValue is a non-required field for NewOrderMultileg.
func (m Message) DiscretionOffsetValue() (*field.DiscretionOffsetValueField, quickfix.MessageRejectError) {
	f := &field.DiscretionOffsetValueField{}
	err := m.Body.Get(f)
	return f, err
}

//GetDiscretionOffsetValue reads a DiscretionOffsetValue from NewOrderMultileg.
func (m Message) GetDiscretionOffsetValue(f *field.DiscretionOffsetValueField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//DiscretionMoveType is a non-required field for NewOrderMultileg.
func (m Message) DiscretionMoveType() (*field.DiscretionMoveTypeField, quickfix.MessageRejectError) {
	f := &field.DiscretionMoveTypeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetDiscretionMoveType reads a DiscretionMoveType from NewOrderMultileg.
func (m Message) GetDiscretionMoveType(f *field.DiscretionMoveTypeField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//DiscretionOffsetType is a non-required field for NewOrderMultileg.
func (m Message) DiscretionOffsetType() (*field.DiscretionOffsetTypeField, quickfix.MessageRejectError) {
	f := &field.DiscretionOffsetTypeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetDiscretionOffsetType reads a DiscretionOffsetType from NewOrderMultileg.
func (m Message) GetDiscretionOffsetType(f *field.DiscretionOffsetTypeField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//DiscretionLimitType is a non-required field for NewOrderMultileg.
func (m Message) DiscretionLimitType() (*field.DiscretionLimitTypeField, quickfix.MessageRejectError) {
	f := &field.DiscretionLimitTypeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetDiscretionLimitType reads a DiscretionLimitType from NewOrderMultileg.
func (m Message) GetDiscretionLimitType(f *field.DiscretionLimitTypeField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//DiscretionRoundDirection is a non-required field for NewOrderMultileg.
func (m Message) DiscretionRoundDirection() (*field.DiscretionRoundDirectionField, quickfix.MessageRejectError) {
	f := &field.DiscretionRoundDirectionField{}
	err := m.Body.Get(f)
	return f, err
}

//GetDiscretionRoundDirection reads a DiscretionRoundDirection from NewOrderMultileg.
func (m Message) GetDiscretionRoundDirection(f *field.DiscretionRoundDirectionField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//DiscretionScope is a non-required field for NewOrderMultileg.
func (m Message) DiscretionScope() (*field.DiscretionScopeField, quickfix.MessageRejectError) {
	f := &field.DiscretionScopeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetDiscretionScope reads a DiscretionScope from NewOrderMultileg.
func (m Message) GetDiscretionScope(f *field.DiscretionScopeField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//TargetStrategy is a non-required field for NewOrderMultileg.
func (m Message) TargetStrategy() (*field.TargetStrategyField, quickfix.MessageRejectError) {
	f := &field.TargetStrategyField{}
	err := m.Body.Get(f)
	return f, err
}

//GetTargetStrategy reads a TargetStrategy from NewOrderMultileg.
func (m Message) GetTargetStrategy(f *field.TargetStrategyField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//TargetStrategyParameters is a non-required field for NewOrderMultileg.
func (m Message) TargetStrategyParameters() (*field.TargetStrategyParametersField, quickfix.MessageRejectError) {
	f := &field.TargetStrategyParametersField{}
	err := m.Body.Get(f)
	return f, err
}

//GetTargetStrategyParameters reads a TargetStrategyParameters from NewOrderMultileg.
func (m Message) GetTargetStrategyParameters(f *field.TargetStrategyParametersField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//ParticipationRate is a non-required field for NewOrderMultileg.
func (m Message) ParticipationRate() (*field.ParticipationRateField, quickfix.MessageRejectError) {
	f := &field.ParticipationRateField{}
	err := m.Body.Get(f)
	return f, err
}

//GetParticipationRate reads a ParticipationRate from NewOrderMultileg.
func (m Message) GetParticipationRate(f *field.ParticipationRateField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//CancellationRights is a non-required field for NewOrderMultileg.
func (m Message) CancellationRights() (*field.CancellationRightsField, quickfix.MessageRejectError) {
	f := &field.CancellationRightsField{}
	err := m.Body.Get(f)
	return f, err
}

//GetCancellationRights reads a CancellationRights from NewOrderMultileg.
func (m Message) GetCancellationRights(f *field.CancellationRightsField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//MoneyLaunderingStatus is a non-required field for NewOrderMultileg.
func (m Message) MoneyLaunderingStatus() (*field.MoneyLaunderingStatusField, quickfix.MessageRejectError) {
	f := &field.MoneyLaunderingStatusField{}
	err := m.Body.Get(f)
	return f, err
}

//GetMoneyLaunderingStatus reads a MoneyLaunderingStatus from NewOrderMultileg.
func (m Message) GetMoneyLaunderingStatus(f *field.MoneyLaunderingStatusField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//RegistID is a non-required field for NewOrderMultileg.
func (m Message) RegistID() (*field.RegistIDField, quickfix.MessageRejectError) {
	f := &field.RegistIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetRegistID reads a RegistID from NewOrderMultileg.
func (m Message) GetRegistID(f *field.RegistIDField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//Designation is a non-required field for NewOrderMultileg.
func (m Message) Designation() (*field.DesignationField, quickfix.MessageRejectError) {
	f := &field.DesignationField{}
	err := m.Body.Get(f)
	return f, err
}

//GetDesignation reads a Designation from NewOrderMultileg.
func (m Message) GetDesignation(f *field.DesignationField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//MultiLegRptTypeReq is a non-required field for NewOrderMultileg.
func (m Message) MultiLegRptTypeReq() (*field.MultiLegRptTypeReqField, quickfix.MessageRejectError) {
	f := &field.MultiLegRptTypeReqField{}
	err := m.Body.Get(f)
	return f, err
}

//GetMultiLegRptTypeReq reads a MultiLegRptTypeReq from NewOrderMultileg.
func (m Message) GetMultiLegRptTypeReq(f *field.MultiLegRptTypeReqField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//NoStrategyParameters is a non-required field for NewOrderMultileg.
func (m Message) NoStrategyParameters() (*field.NoStrategyParametersField, quickfix.MessageRejectError) {
	f := &field.NoStrategyParametersField{}
	err := m.Body.Get(f)
	return f, err
}

//GetNoStrategyParameters reads a NoStrategyParameters from NewOrderMultileg.
func (m Message) GetNoStrategyParameters(f *field.NoStrategyParametersField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//SwapPoints is a non-required field for NewOrderMultileg.
func (m Message) SwapPoints() (*field.SwapPointsField, quickfix.MessageRejectError) {
	f := &field.SwapPointsField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSwapPoints reads a SwapPoints from NewOrderMultileg.
func (m Message) GetSwapPoints(f *field.SwapPointsField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//MatchIncrement is a non-required field for NewOrderMultileg.
func (m Message) MatchIncrement() (*field.MatchIncrementField, quickfix.MessageRejectError) {
	f := &field.MatchIncrementField{}
	err := m.Body.Get(f)
	return f, err
}

//GetMatchIncrement reads a MatchIncrement from NewOrderMultileg.
func (m Message) GetMatchIncrement(f *field.MatchIncrementField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//MaxPriceLevels is a non-required field for NewOrderMultileg.
func (m Message) MaxPriceLevels() (*field.MaxPriceLevelsField, quickfix.MessageRejectError) {
	f := &field.MaxPriceLevelsField{}
	err := m.Body.Get(f)
	return f, err
}

//GetMaxPriceLevels reads a MaxPriceLevels from NewOrderMultileg.
func (m Message) GetMaxPriceLevels(f *field.MaxPriceLevelsField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//SecondaryDisplayQty is a non-required field for NewOrderMultileg.
func (m Message) SecondaryDisplayQty() (*field.SecondaryDisplayQtyField, quickfix.MessageRejectError) {
	f := &field.SecondaryDisplayQtyField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSecondaryDisplayQty reads a SecondaryDisplayQty from NewOrderMultileg.
func (m Message) GetSecondaryDisplayQty(f *field.SecondaryDisplayQtyField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//DisplayWhen is a non-required field for NewOrderMultileg.
func (m Message) DisplayWhen() (*field.DisplayWhenField, quickfix.MessageRejectError) {
	f := &field.DisplayWhenField{}
	err := m.Body.Get(f)
	return f, err
}

//GetDisplayWhen reads a DisplayWhen from NewOrderMultileg.
func (m Message) GetDisplayWhen(f *field.DisplayWhenField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//DisplayMethod is a non-required field for NewOrderMultileg.
func (m Message) DisplayMethod() (*field.DisplayMethodField, quickfix.MessageRejectError) {
	f := &field.DisplayMethodField{}
	err := m.Body.Get(f)
	return f, err
}

//GetDisplayMethod reads a DisplayMethod from NewOrderMultileg.
func (m Message) GetDisplayMethod(f *field.DisplayMethodField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//DisplayLowQty is a non-required field for NewOrderMultileg.
func (m Message) DisplayLowQty() (*field.DisplayLowQtyField, quickfix.MessageRejectError) {
	f := &field.DisplayLowQtyField{}
	err := m.Body.Get(f)
	return f, err
}

//GetDisplayLowQty reads a DisplayLowQty from NewOrderMultileg.
func (m Message) GetDisplayLowQty(f *field.DisplayLowQtyField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//DisplayHighQty is a non-required field for NewOrderMultileg.
func (m Message) DisplayHighQty() (*field.DisplayHighQtyField, quickfix.MessageRejectError) {
	f := &field.DisplayHighQtyField{}
	err := m.Body.Get(f)
	return f, err
}

//GetDisplayHighQty reads a DisplayHighQty from NewOrderMultileg.
func (m Message) GetDisplayHighQty(f *field.DisplayHighQtyField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//DisplayMinIncr is a non-required field for NewOrderMultileg.
func (m Message) DisplayMinIncr() (*field.DisplayMinIncrField, quickfix.MessageRejectError) {
	f := &field.DisplayMinIncrField{}
	err := m.Body.Get(f)
	return f, err
}

//GetDisplayMinIncr reads a DisplayMinIncr from NewOrderMultileg.
func (m Message) GetDisplayMinIncr(f *field.DisplayMinIncrField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//RefreshQty is a non-required field for NewOrderMultileg.
func (m Message) RefreshQty() (*field.RefreshQtyField, quickfix.MessageRejectError) {
	f := &field.RefreshQtyField{}
	err := m.Body.Get(f)
	return f, err
}

//GetRefreshQty reads a RefreshQty from NewOrderMultileg.
func (m Message) GetRefreshQty(f *field.RefreshQtyField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//DisplayQty is a non-required field for NewOrderMultileg.
func (m Message) DisplayQty() (*field.DisplayQtyField, quickfix.MessageRejectError) {
	f := &field.DisplayQtyField{}
	err := m.Body.Get(f)
	return f, err
}

//GetDisplayQty reads a DisplayQty from NewOrderMultileg.
func (m Message) GetDisplayQty(f *field.DisplayQtyField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//PriceProtectionScope is a non-required field for NewOrderMultileg.
func (m Message) PriceProtectionScope() (*field.PriceProtectionScopeField, quickfix.MessageRejectError) {
	f := &field.PriceProtectionScopeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetPriceProtectionScope reads a PriceProtectionScope from NewOrderMultileg.
func (m Message) GetPriceProtectionScope(f *field.PriceProtectionScopeField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//TriggerType is a non-required field for NewOrderMultileg.
func (m Message) TriggerType() (*field.TriggerTypeField, quickfix.MessageRejectError) {
	f := &field.TriggerTypeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetTriggerType reads a TriggerType from NewOrderMultileg.
func (m Message) GetTriggerType(f *field.TriggerTypeField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//TriggerAction is a non-required field for NewOrderMultileg.
func (m Message) TriggerAction() (*field.TriggerActionField, quickfix.MessageRejectError) {
	f := &field.TriggerActionField{}
	err := m.Body.Get(f)
	return f, err
}

//GetTriggerAction reads a TriggerAction from NewOrderMultileg.
func (m Message) GetTriggerAction(f *field.TriggerActionField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//TriggerPrice is a non-required field for NewOrderMultileg.
func (m Message) TriggerPrice() (*field.TriggerPriceField, quickfix.MessageRejectError) {
	f := &field.TriggerPriceField{}
	err := m.Body.Get(f)
	return f, err
}

//GetTriggerPrice reads a TriggerPrice from NewOrderMultileg.
func (m Message) GetTriggerPrice(f *field.TriggerPriceField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//TriggerSymbol is a non-required field for NewOrderMultileg.
func (m Message) TriggerSymbol() (*field.TriggerSymbolField, quickfix.MessageRejectError) {
	f := &field.TriggerSymbolField{}
	err := m.Body.Get(f)
	return f, err
}

//GetTriggerSymbol reads a TriggerSymbol from NewOrderMultileg.
func (m Message) GetTriggerSymbol(f *field.TriggerSymbolField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//TriggerSecurityID is a non-required field for NewOrderMultileg.
func (m Message) TriggerSecurityID() (*field.TriggerSecurityIDField, quickfix.MessageRejectError) {
	f := &field.TriggerSecurityIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetTriggerSecurityID reads a TriggerSecurityID from NewOrderMultileg.
func (m Message) GetTriggerSecurityID(f *field.TriggerSecurityIDField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//TriggerSecurityIDSource is a non-required field for NewOrderMultileg.
func (m Message) TriggerSecurityIDSource() (*field.TriggerSecurityIDSourceField, quickfix.MessageRejectError) {
	f := &field.TriggerSecurityIDSourceField{}
	err := m.Body.Get(f)
	return f, err
}

//GetTriggerSecurityIDSource reads a TriggerSecurityIDSource from NewOrderMultileg.
func (m Message) GetTriggerSecurityIDSource(f *field.TriggerSecurityIDSourceField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//TriggerSecurityDesc is a non-required field for NewOrderMultileg.
func (m Message) TriggerSecurityDesc() (*field.TriggerSecurityDescField, quickfix.MessageRejectError) {
	f := &field.TriggerSecurityDescField{}
	err := m.Body.Get(f)
	return f, err
}

//GetTriggerSecurityDesc reads a TriggerSecurityDesc from NewOrderMultileg.
func (m Message) GetTriggerSecurityDesc(f *field.TriggerSecurityDescField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//TriggerPriceType is a non-required field for NewOrderMultileg.
func (m Message) TriggerPriceType() (*field.TriggerPriceTypeField, quickfix.MessageRejectError) {
	f := &field.TriggerPriceTypeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetTriggerPriceType reads a TriggerPriceType from NewOrderMultileg.
func (m Message) GetTriggerPriceType(f *field.TriggerPriceTypeField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//TriggerPriceTypeScope is a non-required field for NewOrderMultileg.
func (m Message) TriggerPriceTypeScope() (*field.TriggerPriceTypeScopeField, quickfix.MessageRejectError) {
	f := &field.TriggerPriceTypeScopeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetTriggerPriceTypeScope reads a TriggerPriceTypeScope from NewOrderMultileg.
func (m Message) GetTriggerPriceTypeScope(f *field.TriggerPriceTypeScopeField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//TriggerPriceDirection is a non-required field for NewOrderMultileg.
func (m Message) TriggerPriceDirection() (*field.TriggerPriceDirectionField, quickfix.MessageRejectError) {
	f := &field.TriggerPriceDirectionField{}
	err := m.Body.Get(f)
	return f, err
}

//GetTriggerPriceDirection reads a TriggerPriceDirection from NewOrderMultileg.
func (m Message) GetTriggerPriceDirection(f *field.TriggerPriceDirectionField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//TriggerNewPrice is a non-required field for NewOrderMultileg.
func (m Message) TriggerNewPrice() (*field.TriggerNewPriceField, quickfix.MessageRejectError) {
	f := &field.TriggerNewPriceField{}
	err := m.Body.Get(f)
	return f, err
}

//GetTriggerNewPrice reads a TriggerNewPrice from NewOrderMultileg.
func (m Message) GetTriggerNewPrice(f *field.TriggerNewPriceField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//TriggerOrderType is a non-required field for NewOrderMultileg.
func (m Message) TriggerOrderType() (*field.TriggerOrderTypeField, quickfix.MessageRejectError) {
	f := &field.TriggerOrderTypeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetTriggerOrderType reads a TriggerOrderType from NewOrderMultileg.
func (m Message) GetTriggerOrderType(f *field.TriggerOrderTypeField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//TriggerNewQty is a non-required field for NewOrderMultileg.
func (m Message) TriggerNewQty() (*field.TriggerNewQtyField, quickfix.MessageRejectError) {
	f := &field.TriggerNewQtyField{}
	err := m.Body.Get(f)
	return f, err
}

//GetTriggerNewQty reads a TriggerNewQty from NewOrderMultileg.
func (m Message) GetTriggerNewQty(f *field.TriggerNewQtyField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//TriggerTradingSessionID is a non-required field for NewOrderMultileg.
func (m Message) TriggerTradingSessionID() (*field.TriggerTradingSessionIDField, quickfix.MessageRejectError) {
	f := &field.TriggerTradingSessionIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetTriggerTradingSessionID reads a TriggerTradingSessionID from NewOrderMultileg.
func (m Message) GetTriggerTradingSessionID(f *field.TriggerTradingSessionIDField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//TriggerTradingSessionSubID is a non-required field for NewOrderMultileg.
func (m Message) TriggerTradingSessionSubID() (*field.TriggerTradingSessionSubIDField, quickfix.MessageRejectError) {
	f := &field.TriggerTradingSessionSubIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetTriggerTradingSessionSubID reads a TriggerTradingSessionSubID from NewOrderMultileg.
func (m Message) GetTriggerTradingSessionSubID(f *field.TriggerTradingSessionSubIDField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//RefOrderID is a non-required field for NewOrderMultileg.
func (m Message) RefOrderID() (*field.RefOrderIDField, quickfix.MessageRejectError) {
	f := &field.RefOrderIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetRefOrderID reads a RefOrderID from NewOrderMultileg.
func (m Message) GetRefOrderID(f *field.RefOrderIDField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//RefOrderIDSource is a non-required field for NewOrderMultileg.
func (m Message) RefOrderIDSource() (*field.RefOrderIDSourceField, quickfix.MessageRejectError) {
	f := &field.RefOrderIDSourceField{}
	err := m.Body.Get(f)
	return f, err
}

//GetRefOrderIDSource reads a RefOrderIDSource from NewOrderMultileg.
func (m Message) GetRefOrderIDSource(f *field.RefOrderIDSourceField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//PreTradeAnonymity is a non-required field for NewOrderMultileg.
func (m Message) PreTradeAnonymity() (*field.PreTradeAnonymityField, quickfix.MessageRejectError) {
	f := &field.PreTradeAnonymityField{}
	err := m.Body.Get(f)
	return f, err
}

//GetPreTradeAnonymity reads a PreTradeAnonymity from NewOrderMultileg.
func (m Message) GetPreTradeAnonymity(f *field.PreTradeAnonymityField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//ExDestinationIDSource is a non-required field for NewOrderMultileg.
func (m Message) ExDestinationIDSource() (*field.ExDestinationIDSourceField, quickfix.MessageRejectError) {
	f := &field.ExDestinationIDSourceField{}
	err := m.Body.Get(f)
	return f, err
}

//GetExDestinationIDSource reads a ExDestinationIDSource from NewOrderMultileg.
func (m Message) GetExDestinationIDSource(f *field.ExDestinationIDSourceField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//MultilegModel is a non-required field for NewOrderMultileg.
func (m Message) MultilegModel() (*field.MultilegModelField, quickfix.MessageRejectError) {
	f := &field.MultilegModelField{}
	err := m.Body.Get(f)
	return f, err
}

//GetMultilegModel reads a MultilegModel from NewOrderMultileg.
func (m Message) GetMultilegModel(f *field.MultilegModelField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//MultilegPriceMethod is a non-required field for NewOrderMultileg.
func (m Message) MultilegPriceMethod() (*field.MultilegPriceMethodField, quickfix.MessageRejectError) {
	f := &field.MultilegPriceMethodField{}
	err := m.Body.Get(f)
	return f, err
}

//GetMultilegPriceMethod reads a MultilegPriceMethod from NewOrderMultileg.
func (m Message) GetMultilegPriceMethod(f *field.MultilegPriceMethodField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//RiskFreeRate is a non-required field for NewOrderMultileg.
func (m Message) RiskFreeRate() (*field.RiskFreeRateField, quickfix.MessageRejectError) {
	f := &field.RiskFreeRateField{}
	err := m.Body.Get(f)
	return f, err
}

//GetRiskFreeRate reads a RiskFreeRate from NewOrderMultileg.
func (m Message) GetRiskFreeRate(f *field.RiskFreeRateField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//New returns an initialized Message with specified required fields for NewOrderMultileg.
func New(
	clordid *field.ClOrdIDField,
	side *field.SideField,
	nolegs *field.NoLegsField,
	transacttime *field.TransactTimeField,
	ordtype *field.OrdTypeField) Message {
	builder := Message{Message: quickfix.NewMessage()}
	builder.Header.Set(field.NewBeginString(enum.BeginStringFIXT11))
	builder.Header.Set(field.NewDefaultApplVerID(enum.ApplVerID_FIX50SP2))
	builder.Header.Set(field.NewMsgType("AB"))
	builder.Body.Set(clordid)
	builder.Body.Set(side)
	builder.Body.Set(nolegs)
	builder.Body.Set(transacttime)
	builder.Body.Set(ordtype)
	return builder
}

//A RouteOut is the callback type that should be implemented for routing Message
type RouteOut func(msg Message, sessionID quickfix.SessionID) quickfix.MessageRejectError

//Route returns the beginstring, message type, and MessageRoute for this Mesage type
func Route(router RouteOut) (string, string, quickfix.MessageRoute) {
	r := func(msg quickfix.Message, sessionID quickfix.SessionID) quickfix.MessageRejectError {
		return router(Message{msg}, sessionID)
	}
	return enum.ApplVerID_FIX50SP2, "AB", r
}
