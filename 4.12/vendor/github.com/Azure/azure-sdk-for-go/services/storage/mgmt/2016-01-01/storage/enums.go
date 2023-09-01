package storage

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

// AccessTier enumerates the values for access tier.
type AccessTier string

const (
	// Cool ...
	Cool AccessTier = "Cool"
	// Hot ...
	Hot AccessTier = "Hot"
)

// PossibleAccessTierValues returns an array of possible values for the AccessTier const type.
func PossibleAccessTierValues() []AccessTier {
	return []AccessTier{Cool, Hot}
}

// AccountStatus enumerates the values for account status.
type AccountStatus string

const (
	// Available ...
	Available AccountStatus = "Available"
	// Unavailable ...
	Unavailable AccountStatus = "Unavailable"
)

// PossibleAccountStatusValues returns an array of possible values for the AccountStatus const type.
func PossibleAccountStatusValues() []AccountStatus {
	return []AccountStatus{Available, Unavailable}
}

// KeyPermission enumerates the values for key permission.
type KeyPermission string

const (
	// FULL ...
	FULL KeyPermission = "FULL"
	// READ ...
	READ KeyPermission = "READ"
)

// PossibleKeyPermissionValues returns an array of possible values for the KeyPermission const type.
func PossibleKeyPermissionValues() []KeyPermission {
	return []KeyPermission{FULL, READ}
}

// Kind enumerates the values for kind.
type Kind string

const (
	// BlobStorage ...
	BlobStorage Kind = "BlobStorage"
	// Storage ...
	Storage Kind = "Storage"
)

// PossibleKindValues returns an array of possible values for the Kind const type.
func PossibleKindValues() []Kind {
	return []Kind{BlobStorage, Storage}
}

// ProvisioningState enumerates the values for provisioning state.
type ProvisioningState string

const (
	// Creating ...
	Creating ProvisioningState = "Creating"
	// ResolvingDNS ...
	ResolvingDNS ProvisioningState = "ResolvingDNS"
	// Succeeded ...
	Succeeded ProvisioningState = "Succeeded"
)

// PossibleProvisioningStateValues returns an array of possible values for the ProvisioningState const type.
func PossibleProvisioningStateValues() []ProvisioningState {
	return []ProvisioningState{Creating, ResolvingDNS, Succeeded}
}

// Reason enumerates the values for reason.
type Reason string

const (
	// AccountNameInvalid ...
	AccountNameInvalid Reason = "AccountNameInvalid"
	// AlreadyExists ...
	AlreadyExists Reason = "AlreadyExists"
)

// PossibleReasonValues returns an array of possible values for the Reason const type.
func PossibleReasonValues() []Reason {
	return []Reason{AccountNameInvalid, AlreadyExists}
}

// SkuName enumerates the values for sku name.
type SkuName string

const (
	// PremiumLRS ...
	PremiumLRS SkuName = "Premium_LRS"
	// StandardGRS ...
	StandardGRS SkuName = "Standard_GRS"
	// StandardLRS ...
	StandardLRS SkuName = "Standard_LRS"
	// StandardRAGRS ...
	StandardRAGRS SkuName = "Standard_RAGRS"
	// StandardZRS ...
	StandardZRS SkuName = "Standard_ZRS"
)

// PossibleSkuNameValues returns an array of possible values for the SkuName const type.
func PossibleSkuNameValues() []SkuName {
	return []SkuName{PremiumLRS, StandardGRS, StandardLRS, StandardRAGRS, StandardZRS}
}

// SkuTier enumerates the values for sku tier.
type SkuTier string

const (
	// Premium ...
	Premium SkuTier = "Premium"
	// Standard ...
	Standard SkuTier = "Standard"
)

// PossibleSkuTierValues returns an array of possible values for the SkuTier const type.
func PossibleSkuTierValues() []SkuTier {
	return []SkuTier{Premium, Standard}
}

// UsageUnit enumerates the values for usage unit.
type UsageUnit string

const (
	// Bytes ...
	Bytes UsageUnit = "Bytes"
	// BytesPerSecond ...
	BytesPerSecond UsageUnit = "BytesPerSecond"
	// Count ...
	Count UsageUnit = "Count"
	// CountsPerSecond ...
	CountsPerSecond UsageUnit = "CountsPerSecond"
	// Percent ...
	Percent UsageUnit = "Percent"
	// Seconds ...
	Seconds UsageUnit = "Seconds"
)

// PossibleUsageUnitValues returns an array of possible values for the UsageUnit const type.
func PossibleUsageUnitValues() []UsageUnit {
	return []UsageUnit{Bytes, BytesPerSecond, Count, CountsPerSecond, Percent, Seconds}
}