/*
Sets up the global ACRs for SVC and OCP images.
*/
import { getLocationAvailabilityZones } from 'common.bicep'

param ocpAcrName string
param ocpAcrSku string

param svcAcrName string
param svcAcrSku string

param location string

param svcAcrZoneRedundancy string = length(getLocationAvailabilityZones(location)) > 0 ? 'Enabled' : 'Disabled'
param ocpAcrZoneRedundancy string = length(getLocationAvailabilityZones(location)) > 0 ? 'Enabled' : 'Disabled'

module ocpAcr '../modules/acr/acr.bicep' = {
  name: ocpAcrName
  params: {
    acrName: ocpAcrName
    acrSku: ocpAcrSku
    location: location
    zoneRedundancy: svcAcrZoneRedundancy
  }
}

module svcAcr '../modules/acr/acr.bicep' = {
  name: svcAcrSku
  params: {
    acrName: svcAcrName
    acrSku: svcAcrSku
    location: location
    zoneRedundancy: ocpAcrZoneRedundancy
  }
}
