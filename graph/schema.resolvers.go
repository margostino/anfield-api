package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/margostino/owid-api/graph/generated"
	"github.com/margostino/owid-api/graph/model"
)

func (r *queryResolver) O20thCenturyDeathsInUsCdc(ctx context.Context, entity string, year int) (*model.O20thCenturyDeathsInUsCdcDataset, error) {
	var response model.O20thCenturyDeathsInUsCdcDataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) ACenturyOfWorkAndLeisureRameyAndFrancis2009(ctx context.Context, entity string, year int) (*model.ACenturyOfWorkAndLeisureRameyAndFrancis2009Dataset, error) {
	var response model.ACenturyOfWorkAndLeisureRameyAndFrancis2009Dataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) AbsoluteDeathsFromAmbientPm25AirPollutionStateOfGlobalAir(ctx context.Context, entity string, year int) (*model.AbsoluteDeathsFromAmbientPm25AirPollutionStateOfGlobalAirDataset, error) {
	var response model.AbsoluteDeathsFromAmbientPm25AirPollutionStateOfGlobalAirDataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) AbsolutePopulationChangeOwidBasedOnHydeAndUn(ctx context.Context, entity string, year int) (*model.AbsolutePopulationChangeOwidBasedOnHydeAndUnDataset, error) {
	var response model.AbsolutePopulationChangeOwidBasedOnHydeAndUnDataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) AccessToFinancialAccountOrServicesPercWorldBank2014(ctx context.Context, entity string, year int) (*model.AccessToFinancialAccountOrServicesPercWorldBank2014Dataset, error) {
	var response model.AccessToFinancialAccountOrServicesPercWorldBank2014Dataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) AdjustedNetSavingsPerCapitaWorldBankWdi2017(ctx context.Context, entity string, year int) (*model.AdjustedNetSavingsPerCapitaWorldBankWdi2017Dataset, error) {
	var response model.AdjustedNetSavingsPerCapitaWorldBankWdi2017Dataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) AdultLiteracyProficiencyWorldBankEdstatsAndStepSkillsMeasurementProgram(ctx context.Context, entity string, year int) (*model.AdultLiteracyProficiencyWorldBankEdstatsAndStepSkillsMeasurementProgramDataset, error) {
	var response model.AdultLiteracyProficiencyWorldBankEdstatsAndStepSkillsMeasurementProgramDataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) AdultObesityByRegionFao2017(ctx context.Context, entity string, year int) (*model.AdultObesityByRegionFao2017Dataset, error) {
	var response model.AdultObesityByRegionFao2017Dataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) AffordabilityOfDietsSofi2021(ctx context.Context, entity string, year int) (*model.AffordabilityOfDietsSofi2021Dataset, error) {
	var response model.AffordabilityOfDietsSofi2021Dataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) AgriculturalPolicySupportAgrimonitor2017(ctx context.Context, entity string, year int) (*model.AgriculturalPolicySupportAgrimonitor2017Dataset, error) {
	var response model.AgriculturalPolicySupportAgrimonitor2017Dataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) AgriculturalTotalFactorProductivityUsda(ctx context.Context, entity string, year int) (*model.AgriculturalTotalFactorProductivityUsdaDataset, error) {
	var response model.AgriculturalTotalFactorProductivityUsdaDataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) AgricultureInEngland1270_1870BankOfEngland2017(ctx context.Context, entity string, year int) (*model.AgricultureInEngland12701870BankOfEngland2017Dataset, error) {
	var response model.AgricultureInEngland12701870BankOfEngland2017Dataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) AirPollutantEmissionsOecd(ctx context.Context, entity string, year int) (*model.AirPollutantEmissionsOecdDataset, error) {
	var response model.AirPollutantEmissionsOecdDataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) AirPollutionByCityFouquetAndDpcc2011(ctx context.Context, entity string, year int) (*model.AirPollutionByCityFouquetAndDpcc2011Dataset, error) {
	var response model.AirPollutionByCityFouquetAndDpcc2011Dataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) AirPollutionSourcesInTheUkDefra(ctx context.Context, entity string, year int) (*model.AirPollutionSourcesInTheUkDefraDataset, error) {
	var response model.AirPollutionSourcesInTheUkDefraDataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) AirTravelTripsPerCapitaAirbus2019(ctx context.Context, entity string, year int) (*model.AirTravelTripsPerCapitaAirbus2019Dataset, error) {
	var response model.AirTravelTripsPerCapitaAirbus2019Dataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) AirlineHijackingAviationSafetyNetwork(ctx context.Context, entity string, year int) (*model.AirlineHijackingAviationSafetyNetworkDataset, error) {
	var response model.AirlineHijackingAviationSafetyNetworkDataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) AlcoholConsumptionByTypeSince1890AlexanderAndHolmes2017(ctx context.Context, entity string, year int) (*model.AlcoholConsumptionByTypeSince1890AlexanderAndHolmes2017Dataset, error) {
	var response model.AlcoholConsumptionByTypeSince1890AlexanderAndHolmes2017Dataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) AlcoholConsumptionInUsaSince1850Niaaa(ctx context.Context, entity string, year int) (*model.AlcoholConsumptionInUsaSince1850NiaaaDataset, error) {
	var response model.AlcoholConsumptionInUsaSince1850NiaaaDataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) AlcoholConsumptionSince1890AlexanderAndHolmes2017(ctx context.Context, entity string, year int) (*model.AlcoholConsumptionSince1890AlexanderAndHolmes2017Dataset, error) {
	var response model.AlcoholConsumptionSince1890AlexanderAndHolmes2017Dataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) AlcoholExpenditureInTheUsaLongTermUsda2018(ctx context.Context, entity string, year int) (*model.AlcoholExpenditureInTheUsaLongTermUsda2018Dataset, error) {
	var response model.AlcoholExpenditureInTheUsaLongTermUsda2018Dataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) AnnualShareOfCo2EmissionsOwidBasedOnGcp2017(ctx context.Context, entity string, year int) (*model.AnnualShareOfCo2EmissionsOwidBasedOnGcp2017Dataset, error) {
	var response model.AnnualShareOfCo2EmissionsOwidBasedOnGcp2017Dataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) AnnualWorldPopulationGrowthRateOwid(ctx context.Context, entity string, year int) (*model.AnnualWorldPopulationGrowthRateOwidDataset, error) {
	var response model.AnnualWorldPopulationGrowthRateOwidDataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) AntibioticUseInLivestockEuropeanCommissionAndVanBoeckelEtAl(ctx context.Context, entity string, year int) (*model.AntibioticUseInLivestockEuropeanCommissionAndVanBoeckelEtAlDataset, error) {
	var response model.AntibioticUseInLivestockEuropeanCommissionAndVanBoeckelEtAlDataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) AntibioticUseInLivestock2030BoeckelEtAl2017(ctx context.Context, entity string, year int) (*model.AntibioticUseInLivestock2030BoeckelEtAl2017Dataset, error) {
	var response model.AntibioticUseInLivestock2030BoeckelEtAl2017Dataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) ArableLandPerCropOutputPinFao2019(ctx context.Context, entity string, year int) (*model.ArableLandPerCropOutputPinFao2019Dataset, error) {
	var response model.ArableLandPerCropOutputPinFao2019Dataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) ArchaeologicalLandUseStephensEtAl2019(ctx context.Context, entity string, year int) (*model.ArchaeologicalLandUseStephensEtAl2019Dataset, error) {
	var response model.ArchaeologicalLandUseStephensEtAl2019Dataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) ArmedForcesPersonnelAsAShareOfTheTotalPopulationOwidBasedOnWorldBank(ctx context.Context, entity string, year int) (*model.ArmedForcesPersonnelAsAShareOfTheTotalPopulationOwidBasedOnWorldBankDataset, error) {
	var response model.ArmedForcesPersonnelAsAShareOfTheTotalPopulationOwidBasedOnWorldBankDataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) AttainableYieldsMuellerEtAl2012(ctx context.Context, entity string, year int) (*model.AttainableYieldsMuellerEtAl2012Dataset, error) {
	var response model.AttainableYieldsMuellerEtAl2012Dataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) AttitudesToVaccinesWellcomeTrust2019(ctx context.Context, entity string, year int) (*model.AttitudesToVaccinesWellcomeTrust2019Dataset, error) {
	var response model.AttitudesToVaccinesWellcomeTrust2019Dataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) AverageHarmonisedLearningOutcomeScore2005_2015AltinokAngristAndPatrinos2018(ctx context.Context, entity string, year int) (*model.AverageHarmonisedLearningOutcomeScore20052015AltinokAngristAndPatrinos2018Dataset, error) {
	var response model.AverageHarmonisedLearningOutcomeScore20052015AltinokAngristAndPatrinos2018Dataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) AverageMonthlyIncomesOrConsumptionByDecileAndQuintilePovcalnet2019(ctx context.Context, entity string, year int) (*model.AverageMonthlyIncomesOrConsumptionByDecileAndQuintilePovcalnet2019Dataset, error) {
	var response model.AverageMonthlyIncomesOrConsumptionByDecileAndQuintilePovcalnet2019Dataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) AviationAccidentsAndFatalitiesByFlightPhaseAsn2019(ctx context.Context, entity string, year int) (*model.AviationAccidentsAndFatalitiesByFlightPhaseAsn2019Dataset, error) {
	var response model.AviationAccidentsAndFatalitiesByFlightPhaseAsn2019Dataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) AviationPassengerKilometresAndCo2EmissionsIcct(ctx context.Context, entity string, year int) (*model.AviationPassengerKilometresAndCo2EmissionsIcctDataset, error) {
	var response model.AviationPassengerKilometresAndCo2EmissionsIcctDataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) BasicReadingAndMathsSkillsWorldDevelopmentReport2018(ctx context.Context, entity string, year int) (*model.BasicReadingAndMathsSkillsWorldDevelopmentReport2018Dataset, error) {
	var response model.BasicReadingAndMathsSkillsWorldDevelopmentReport2018Dataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) BiodiversityHabitatLossWilliamsEtAl2021(ctx context.Context, entity string, year int) (*model.BiodiversityHabitatLossWilliamsEtAl2021Dataset, error) {
	var response model.BiodiversityHabitatLossWilliamsEtAl2021Dataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) BiomassAndTaxaAbundanceBarOnEtAl2018(ctx context.Context, entity string, year int) (*model.BiomassAndTaxaAbundanceBarOnEtAl2018Dataset, error) {
	var response model.BiomassAndTaxaAbundanceBarOnEtAl2018Dataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) BirthsOutsideOfMarriage(ctx context.Context, entity string, year int) (*model.BirthsOutsideOfMarriageDataset, error) {
	var response model.BirthsOutsideOfMarriageDataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) BooksBuringhAndVanZanden2009(ctx context.Context, entity string, year int) (*model.BooksBuringhAndVanZanden2009Dataset, error) {
	var response model.BooksBuringhAndVanZanden2009Dataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) BourguignonAndMorrison2002AndWorldBankPovcalnet2015(ctx context.Context, entity string, year int) (*model.BourguignonAndMorrison2002AndWorldBankPovcalnet2015Dataset, error) {
	var response model.BourguignonAndMorrison2002AndWorldBankPovcalnet2015Dataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) Cfc11ExpectedAndMeasuredConcentrationsMontzkaEtAl2018(ctx context.Context, entity string, year int) (*model.Cfc11ExpectedAndMeasuredConcentrationsMontzkaEtAl2018Dataset, error) {
	var response model.Cfc11ExpectedAndMeasuredConcentrationsMontzkaEtAl2018Dataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) Cfc11ExpectedAndMeasuredRateOfChangeMontzkaEtAl2018(ctx context.Context, entity string, year int) (*model.Cfc11ExpectedAndMeasuredRateOfChangeMontzkaEtAl2018Dataset, error) {
	var response model.Cfc11ExpectedAndMeasuredRateOfChangeMontzkaEtAl2018Dataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) Co2FootprintBreakdownPerCapitaGoodall2011(ctx context.Context, entity string, year int) (*model.Co2FootprintBreakdownPerCapitaGoodall2011Dataset, error) {
	var response model.Co2FootprintBreakdownPerCapitaGoodall2011Dataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) Co2EmissionFactorsIpcc2006(ctx context.Context, entity string, year int) (*model.Co2EmissionFactorsIpcc2006Dataset, error) {
	var response model.Co2EmissionFactorsIpcc2006Dataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) Co2EmissionsByCityC40Cities2018(ctx context.Context, entity string, year int) (*model.Co2EmissionsByCityC40Cities2018Dataset, error) {
	var response model.Co2EmissionsByCityC40Cities2018Dataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) Co2EmissionsBySectorCait2020(ctx context.Context, entity string, year int) (*model.Co2EmissionsBySectorCait2020Dataset, error) {
	var response model.Co2EmissionsBySectorCait2020Dataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) Co2EmissionsBySectorCait2021(ctx context.Context, entity string, year int) (*model.Co2EmissionsBySectorCait2021Dataset, error) {
	var response model.Co2EmissionsBySectorCait2021Dataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) Co2EmissionsBySourceCdiac2016(ctx context.Context, entity string, year int) (*model.Co2EmissionsBySourceCdiac2016Dataset, error) {
	var response model.Co2EmissionsBySourceCdiac2016Dataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) Co2EmissionsInTradeAsPercProductionGlobalCarbonProject2014(ctx context.Context, entity string, year int) (*model.Co2EmissionsInTradeAsPercProductionGlobalCarbonProject2014Dataset, error) {
	var response model.Co2EmissionsInTradeAsPercProductionGlobalCarbonProject2014Dataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) Co2FromCementCdiac2017(ctx context.Context, entity string, year int) (*model.Co2FromCementCdiac2017Dataset, error) {
	var response model.Co2FromCementCdiac2017Dataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) Co2FromFlaringCdiac2017(ctx context.Context, entity string, year int) (*model.Co2FromFlaringCdiac2017Dataset, error) {
	var response model.Co2FromFlaringCdiac2017Dataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) Co2FromGasCdiac2017(ctx context.Context, entity string, year int) (*model.Co2FromGasCdiac2017Dataset, error) {
	var response model.Co2FromGasCdiac2017Dataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) Co2FromLiquidCdiac2017(ctx context.Context, entity string, year int) (*model.Co2FromLiquidCdiac2017Dataset, error) {
	var response model.Co2FromLiquidCdiac2017Dataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) Co2FromSolidFuelCdiac2017(ctx context.Context, entity string, year int) (*model.Co2FromSolidFuelCdiac2017Dataset, error) {
	var response model.Co2FromSolidFuelCdiac2017Dataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) Co2IntensityOfTransportByModeUkBeis(ctx context.Context, entity string, year int) (*model.Co2IntensityOfTransportByModeUkBeisDataset, error) {
	var response model.Co2IntensityOfTransportByModeUkBeisDataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) Co2MitigationCurvesFor15cAndrewsAndGcp2019(ctx context.Context, entity string, year int) (*model.Co2MitigationCurvesFor15cAndrewsAndGcp2019Dataset, error) {
	var response model.Co2MitigationCurvesFor15cAndrewsAndGcp2019Dataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) Co2MitigationCurvesFor2cAndrewsAndGcp2019(ctx context.Context, entity string, year int) (*model.Co2MitigationCurvesFor2cAndrewsAndGcp2019Dataset, error) {
	var response model.Co2MitigationCurvesFor2cAndrewsAndGcp2019Dataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) Co2PerYearByRegionCdiac2017(ctx context.Context, entity string, year int) (*model.Co2PerYearByRegionCdiac2017Dataset, error) {
	var response model.Co2PerYearByRegionCdiac2017Dataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) Co2GdpCouplingOwidBasedOnWorldBank(ctx context.Context, entity string, year int) (*model.Co2GdpCouplingOwidBasedOnWorldBankDataset, error) {
	var response model.Co2GdpCouplingOwidBasedOnWorldBankDataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) CovidGovernmentResponseOxbsg(ctx context.Context, entity string, year int) (*model.CovidGovernmentResponseOxbsgDataset, error) {
	var response model.CovidGovernmentResponseOxbsgDataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) CovidTestingTimeSeriesData(ctx context.Context, entity string, year int) (*model.CovidTestingTimeSeriesDataDataset, error) {
	var response model.CovidTestingTimeSeriesDataDataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) Covid2019Ecdc2020(ctx context.Context, entity string, year int) (*model.Covid2019Ecdc2020Dataset, error) {
	var response model.Covid2019Ecdc2020Dataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) Covid2019HospitalAndIcu(ctx context.Context, entity string, year int) (*model.Covid2019HospitalAndIcuDataset, error) {
	var response model.Covid2019HospitalAndIcuDataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) CaloriesLostByFoodGroupAndRegionWri2013(ctx context.Context, entity string, year int) (*model.CaloriesLostByFoodGroupAndRegionWri2013Dataset, error) {
	var response model.CaloriesLostByFoodGroupAndRegionWri2013Dataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) CancerDeathRatesInTheUsOverTheLongTermAmericanCancerSociety(ctx context.Context, entity string, year int) (*model.CancerDeathRatesInTheUsOverTheLongTermAmericanCancerSocietyDataset, error) {
	var response model.CancerDeathRatesInTheUsOverTheLongTermAmericanCancerSocietyDataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) CancerDeathsGroupedOwidBasedOnIhme(ctx context.Context, entity string, year int) (*model.CancerDeathsGroupedOwidBasedOnIhmeDataset, error) {
	var response model.CancerDeathsGroupedOwidBasedOnIhmeDataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) CapitalCityPopulationUnUrbanizationProspects2018(ctx context.Context, entity string, year int) (*model.CapitalCityPopulationUnUrbanizationProspects2018Dataset, error) {
	var response model.CapitalCityPopulationUnUrbanizationProspects2018Dataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) CarbonIntensityKgco2moneyMadissonWorldBankCdiac(ctx context.Context, entity string, year int) (*model.CarbonIntensityKgco2moneyMadissonWorldBankCdiacDataset, error) {
	var response model.CarbonIntensityKgco2moneyMadissonWorldBankCdiacDataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) CausesOfChildMortalityIhmeGlobalBurdenOfDiseaseStudy2017(ctx context.Context, entity string, year int) (*model.CausesOfChildMortalityIhmeGlobalBurdenOfDiseaseStudy2017Dataset, error) {
	var response model.CausesOfChildMortalityIhmeGlobalBurdenOfDiseaseStudy2017Dataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) CausesOfDeathVsMediaCoverageShenEtAl2018(ctx context.Context, entity string, year int) (*model.CausesOfDeathVsMediaCoverageShenEtAl2018Dataset, error) {
	var response model.CausesOfDeathVsMediaCoverageShenEtAl2018Dataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) CausesOfDeathVsMediaCoverageShen(ctx context.Context, entity string, year int) (*model.CausesOfDeathVsMediaCoverageShenDataset, error) {
	var response model.CausesOfDeathVsMediaCoverageShenDataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) CausesOfInfantDeathInBoysAndGirlsIhme2018(ctx context.Context, entity string, year int) (*model.CausesOfInfantDeathInBoysAndGirlsIhme2018Dataset, error) {
	var response model.CausesOfInfantDeathInBoysAndGirlsIhme2018Dataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) CerealAllocationToFoodFeedFuelOwidBasedOnFao(ctx context.Context, entity string, year int) (*model.CerealAllocationToFoodFeedFuelOwidBasedOnFaoDataset, error) {
	var response model.CerealAllocationToFoodFeedFuelOwidBasedOnFaoDataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) CerealYieldIndexWorldBank2017AndOwid(ctx context.Context, entity string, year int) (*model.CerealYieldIndexWorldBank2017AndOwidDataset, error) {
	var response model.CerealYieldIndexWorldBank2017AndOwidDataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) ChangeInGlobalHungerIndex1992_2017Listed2017GlobalHungerIndex2017(ctx context.Context, entity string, year int) (*model.ChangeInGlobalHungerIndex19922017Listed2017GlobalHungerIndex2017Dataset, error) {
	var response model.ChangeInGlobalHungerIndex19922017Listed2017GlobalHungerIndex2017Dataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) ChartbookOfEconomicInequalityGini2016(ctx context.Context, entity string, year int) (*model.ChartbookOfEconomicInequalityGini2016Dataset, error) {
	var response model.ChartbookOfEconomicInequalityGini2016Dataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) ChildLaborWorld1950_1995Basu1999(ctx context.Context, entity string, year int) (*model.ChildLaborWorld19501995Basu1999Dataset, error) {
	var response model.ChildLaborWorld19501995Basu1999Dataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) ChildLaborWorldIloIlo2017(ctx context.Context, entity string, year int) (*model.ChildLaborWorldIloIlo2017Dataset, error) {
	var response model.ChildLaborWorldIloIlo2017Dataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) ChildLaborItalyHistoricTonioloGAndVecchiG2007(ctx context.Context, entity string, year int) (*model.ChildLaborItalyHistoricTonioloGAndVecchiG2007Dataset, error) {
	var response model.ChildLaborItalyHistoricTonioloGAndVecchiG2007Dataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) ChildLaborUkHistoricCunninghamHAndViazzoPp1996(ctx context.Context, entity string, year int) (*model.ChildLaborUkHistoricCunninghamHAndViazzoPp1996Dataset, error) {
	var response model.ChildLaborUkHistoricCunninghamHAndViazzoPp1996Dataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) ChildLaborInUsEconomicHistoryAssociation2017(ctx context.Context, entity string, year int) (*model.ChildLaborInUsEconomicHistoryAssociation2017Dataset, error) {
	var response model.ChildLaborInUsEconomicHistoryAssociation2017Dataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) ChildLaborInUsLong1958(ctx context.Context, entity string, year int) (*model.ChildLaborInUsLong1958Dataset, error) {
	var response model.ChildLaborInUsLong1958Dataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) ChildMarriageUnicef2017(ctx context.Context, entity string, year int) (*model.ChildMarriageUnicef2017Dataset, error) {
	var response model.ChildMarriageUnicef2017Dataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) ChildMortalityEstimatesCmeInfo2018(ctx context.Context, entity string, year int) (*model.ChildMortalityEstimatesCmeInfo2018Dataset, error) {
	var response model.ChildMortalityEstimatesCmeInfo2018Dataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) ChildMortalityRatesCompleteGapminderV10_2017(ctx context.Context, entity string, year int) (*model.ChildMortalityRatesCompleteGapminderV102017Dataset, error) {
	var response model.ChildMortalityRatesCompleteGapminderV102017Dataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) ChildMortalityRatesSelectedGapminderV10_2017(ctx context.Context, entity string, year int) (*model.ChildMortalityRatesSelectedGapminderV102017Dataset, error) {
	var response model.ChildMortalityRatesSelectedGapminderV102017Dataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) ChildDeathsByLifeStageOwidBasedOnUnIgme(ctx context.Context, entity string, year int) (*model.ChildDeathsByLifeStageOwidBasedOnUnIgmeDataset, error) {
	var response model.ChildDeathsByLifeStageOwidBasedOnUnIgmeDataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) ChildDeathsUnitedNationsPopulationDivision2015(ctx context.Context, entity string, year int) (*model.ChildDeathsUnitedNationsPopulationDivision2015Dataset, error) {
	var response model.ChildDeathsUnitedNationsPopulationDivision2015Dataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) ChildMortalityGapminder2013(ctx context.Context, entity string, year int) (*model.ChildMortalityGapminder2013Dataset, error) {
	var response model.ChildMortalityGapminder2013Dataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) ChildMortalityByIncomeLevel1960_2012WorldBankWdi2016(ctx context.Context, entity string, year int) (*model.ChildMortalityByIncomeLevel19602012WorldBankWdi2016Dataset, error) {
	var response model.ChildMortalityByIncomeLevel19602012WorldBankWdi2016Dataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) ChildMortalityDataIhme2017(ctx context.Context, entity string, year int) (*model.ChildMortalityDataIhme2017Dataset, error) {
	var response model.ChildMortalityDataIhme2017Dataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) ChildMortalityEstimatesGapminder2015(ctx context.Context, entity string, year int) (*model.ChildMortalityEstimatesGapminder2015Dataset, error) {
	var response model.ChildMortalityEstimatesGapminder2015Dataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) ChildMortality1950_2017Ihme2017(ctx context.Context, entity string, year int) (*model.ChildMortality19502017Ihme2017Dataset, error) {
	var response model.ChildMortality19502017Ihme2017Dataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) ChildViolenceEndingViolenceInChildhoodReport2017(ctx context.Context, entity string, year int) (*model.ChildViolenceEndingViolenceInChildhoodReport2017Dataset, error) {
	var response model.ChildViolenceEndingViolenceInChildhoodReport2017Dataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) ChildrenThatDiedBefore5YearsOfAgePerWomanGapminder2017(ctx context.Context, entity string, year int) (*model.ChildrenThatDiedBefore5YearsOfAgePerWomanGapminder2017Dataset, error) {
	var response model.ChildrenThatDiedBefore5YearsOfAgePerWomanGapminder2017Dataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) ChildrenThatSurvivedPast5YearsOfAgePerWomanGapminder2017(ctx context.Context, entity string, year int) (*model.ChildrenThatSurvivedPast5YearsOfAgePerWomanGapminder2017Dataset, error) {
	var response model.ChildrenThatSurvivedPast5YearsOfAgePerWomanGapminder2017Dataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) ChinaShareOfWorldPovertyWorldBankWdi2017(ctx context.Context, entity string, year int) (*model.ChinaShareOfWorldPovertyWorldBankWdi2017Dataset, error) {
	var response model.ChinaShareOfWorldPovertyWorldBankWdi2017Dataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) CityPopulations1950_2035UnUrbanizationProspects2018(ctx context.Context, entity string, year int) (*model.CityPopulations19502035UnUrbanizationProspects2018Dataset, error) {
	var response model.CityPopulations19502035UnUrbanizationProspects2018Dataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) ClarkFlecheAndSenikHappinessInequality(ctx context.Context, entity string, year int) (*model.ClarkFlecheAndSenikHappinessInequalityDataset, error) {
	var response model.ClarkFlecheAndSenikHappinessInequalityDataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) CoalOutputAndEmploymentInUkBeis2020(ctx context.Context, entity string, year int) (*model.CoalOutputAndEmploymentInUkBeis2020Dataset, error) {
	var response model.CoalOutputAndEmploymentInUkBeis2020Dataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) CoalProductionTheShiftProject(ctx context.Context, entity string, year int) (*model.CoalProductionTheShiftProjectDataset, error) {
	var response model.CoalProductionTheShiftProjectDataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) CoefficientOfVariationCvInCaloricIntake(ctx context.Context, entity string, year int) (*model.CoefficientOfVariationCvInCaloricIntakeDataset, error) {
	var response model.CoefficientOfVariationCvInCaloricIntakeDataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) ConflictAndTerrorismDeathsOwidBasedOnIhmeAndGtd(ctx context.Context, entity string, year int) (*model.ConflictAndTerrorismDeathsOwidBasedOnIhmeAndGtdDataset, error) {
	var response model.ConflictAndTerrorismDeathsOwidBasedOnIhmeAndGtdDataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) ConflictDeathsByCountryUcdp2019(ctx context.Context, entity string, year int) (*model.ConflictDeathsByCountryUcdp2019Dataset, error) {
	var response model.ConflictDeathsByCountryUcdp2019Dataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) ConflictDeathsUcdpGeoreferencedEventData2019(ctx context.Context, entity string, year int) (*model.ConflictDeathsUcdpGeoreferencedEventData2019Dataset, error) {
	var response model.ConflictDeathsUcdpGeoreferencedEventData2019Dataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) ConsumerExpenditureOnFoodUsda2017(ctx context.Context, entity string, year int) (*model.ConsumerExpenditureOnFoodUsda2017Dataset, error) {
	var response model.ConsumerExpenditureOnFoodUsda2017Dataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) ConsumptionSharesInSelectedNonEssentialProductsWorldBankGlobalConsumptionDatabase(ctx context.Context, entity string, year int) (*model.ConsumptionSharesInSelectedNonEssentialProductsWorldBankGlobalConsumptionDatabaseDataset, error) {
	var response model.ConsumptionSharesInSelectedNonEssentialProductsWorldBankGlobalConsumptionDatabaseDataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) ConsumptionVsProductionBasedCo2EmissionsSharesBasedOnGcpAndUn(ctx context.Context, entity string, year int) (*model.ConsumptionVsProductionBasedCo2EmissionsSharesBasedOnGcpAndUnDataset, error) {
	var response model.ConsumptionVsProductionBasedCo2EmissionsSharesBasedOnGcpAndUnDataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) CorporalPunishmentInSchoolsLongitudinalEvidenceFromEthiopiaIndiaPeruAndVietnamUnicef2015(ctx context.Context, entity string, year int) (*model.CorporalPunishmentInSchoolsLongitudinalEvidenceFromEthiopiaIndiaPeruAndVietnamUnicef2015Dataset, error) {
	var response model.CorporalPunishmentInSchoolsLongitudinalEvidenceFromEthiopiaIndiaPeruAndVietnamUnicef2015Dataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) CorrelatesOfWarNationalMaterialCapabilitiesV40(ctx context.Context, entity string, year int) (*model.CorrelatesOfWarNationalMaterialCapabilitiesV40Dataset, error) {
	var response model.CorrelatesOfWarNationalMaterialCapabilitiesV40Dataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) CorruptionPerceptionIndexTransparencyInternational2018(ctx context.Context, entity string, year int) (*model.CorruptionPerceptionIndexTransparencyInternational2018Dataset, error) {
	var response model.CorruptionPerceptionIndexTransparencyInternational2018Dataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) CountriesContinents(ctx context.Context, entity string, year int) (*model.CountriesContinentsDataset, error) {
	var response model.CountriesContinentsDataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) CountryIncomeClassificationWorldBank2017(ctx context.Context, entity string, year int) (*model.CountryIncomeClassificationWorldBank2017Dataset, error) {
	var response model.CountryIncomeClassificationWorldBank2017Dataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) CountryLevelLandPrecipitationDelaware(ctx context.Context, entity string, year int) (*model.CountryLevelLandPrecipitationDelawareDataset, error) {
	var response model.CountryLevelLandPrecipitationDelawareDataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) CrossCountryLiteracyRatesWorldBankCiaWorldFactbookAndOtherSources(ctx context.Context, entity string, year int) (*model.CrossCountryLiteracyRatesWorldBankCiaWorldFactbookAndOtherSourcesDataset, error) {
	var response model.CrossCountryLiteracyRatesWorldBankCiaWorldFactbookAndOtherSourcesDataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) CrudeBirthAndDeathRatesPer1_000EnglandAndWales1541_2015WrigleyAndSchofieldMitchellUkOns(ctx context.Context, entity string, year int) (*model.CrudeBirthAndDeathRatesPer1000EnglandAndWales15412015WrigleyAndSchofieldMitchellUkOnsDataset, error) {
	var response model.CrudeBirthAndDeathRatesPer1000EnglandAndWales15412015WrigleyAndSchofieldMitchellUkOnsDataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) CrudeOilConsumptionVsPriceBpStatistics2016(ctx context.Context, entity string, year int) (*model.CrudeOilConsumptionVsPriceBpStatistics2016Dataset, error) {
	var response model.CrudeOilConsumptionVsPriceBpStatistics2016Dataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) CrudeMarriageRateOwidBasedOnUnOecdEurostatAndOtherSources(ctx context.Context, entity string, year int) (*model.CrudeMarriageRateOwidBasedOnUnOecdEurostatAndOtherSourcesDataset, error) {
	var response model.CrudeMarriageRateOwidBasedOnUnOecdEurostatAndOtherSourcesDataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) CumulativeCo2EmissionsShareOwidBasedOnGcp2017(ctx context.Context, entity string, year int) (*model.CumulativeCo2EmissionsShareOwidBasedOnGcp2017Dataset, error) {
	var response model.CumulativeCo2EmissionsShareOwidBasedOnGcp2017Dataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) CumulativeShareOfMarriagesEndingInDivorceEnglandAndWalesUkOns(ctx context.Context, entity string, year int) (*model.CumulativeShareOfMarriagesEndingInDivorceEnglandAndWalesUkOnsDataset, error) {
	var response model.CumulativeShareOfMarriagesEndingInDivorceEnglandAndWalesUkOnsDataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) CurrentGdpBritishPoundsFouquinAndHugotCepii2016(ctx context.Context, entity string, year int) (*model.CurrentGdpBritishPoundsFouquinAndHugotCepii2016Dataset, error) {
	var response model.CurrentGdpBritishPoundsFouquinAndHugotCepii2016Dataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) D1VsD10D1IncomeconsumptionPovcal2018(ctx context.Context, entity string, year int) (*model.D1VsD10D1IncomeconsumptionPovcal2018Dataset, error) {
	var response model.D1VsD10D1IncomeconsumptionPovcal2018Dataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) DailyFatSupplyFao2017(ctx context.Context, entity string, year int) (*model.DailyFatSupplyFao2017Dataset, error) {
	var response model.DailyFatSupplyFao2017Dataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) DailyProteinSupplyFao2017(ctx context.Context, entity string, year int) (*model.DailyProteinSupplyFao2017Dataset, error) {
	var response model.DailyProteinSupplyFao2017Dataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) DailySupplyOfCaloriesPerPersonOwidBasedOnUnFaoAndHistoricalSources(ctx context.Context, entity string, year int) (*model.DailySupplyOfCaloriesPerPersonOwidBasedOnUnFaoAndHistoricalSourcesDataset, error) {
	var response model.DailySupplyOfCaloriesPerPersonOwidBasedOnUnFaoAndHistoricalSourcesDataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) DaysAndHoursOfWorkInOldAndNewWorldsHubermanAndMinns2007(ctx context.Context, entity string, year int) (*model.DaysAndHoursOfWorkInOldAndNewWorldsHubermanAndMinns2007Dataset, error) {
	var response model.DaysAndHoursOfWorkInOldAndNewWorldsHubermanAndMinns2007Dataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) DeathRateByAgeGroupInEnglandAndWalesOns(ctx context.Context, entity string, year int) (*model.DeathRateByAgeGroupInEnglandAndWalesOnsDataset, error) {
	var response model.DeathRateByAgeGroupInEnglandAndWalesOnsDataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) DeathsAttributedToAirPollutionLelieveldEtAl2019(ctx context.Context, entity string, year int) (*model.DeathsAttributedToAirPollutionLelieveldEtAl2019Dataset, error) {
	var response model.DeathsAttributedToAirPollutionLelieveldEtAl2019Dataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) DeathsByWorldRegionWho2016(ctx context.Context, entity string, year int) (*model.DeathsByWorldRegionWho2016Dataset, error) {
	var response model.DeathsByWorldRegionWho2016Dataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) DeathsFromChernobylRangeOfLongTermEstimatesWho2005FairlieAndSumner2006CardisEtAl2006(ctx context.Context, entity string, year int) (*model.DeathsFromChernobylRangeOfLongTermEstimatesWho2005FairlieAndSumner2006CardisEtAl2006Dataset, error) {
	var response model.DeathsFromChernobylRangeOfLongTermEstimatesWho2005FairlieAndSumner2006CardisEtAl2006Dataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) DeathsFromFukushimaWho2013_2016(ctx context.Context, entity string, year int) (*model.DeathsFromFukushimaWho20132016Dataset, error) {
	var response model.DeathsFromFukushimaWho20132016Dataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) DeathsFromSmalllpoxAndAllCausesInLondon1629_1902(ctx context.Context, entity string, year int) (*model.DeathsFromSmalllpoxAndAllCausesInLondon16291902Dataset, error) {
	var response model.DeathsFromSmalllpoxAndAllCausesInLondon16291902Dataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) DeathsFromSmallpoxPerMillionPopulationEdwardes1902(ctx context.Context, entity string, year int) (*model.DeathsFromSmallpoxPerMillionPopulationEdwardes1902Dataset, error) {
	var response model.DeathsFromSmallpoxPerMillionPopulationEdwardes1902Dataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) DeathsPerTwhEnergyProductionMarkandyaAndWilkinsonSovacoolEtAl(ctx context.Context, entity string, year int) (*model.DeathsPerTwhEnergyProductionMarkandyaAndWilkinsonSovacoolEtAlDataset, error) {
	var response model.DeathsPerTwhEnergyProductionMarkandyaAndWilkinsonSovacoolEtAlDataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) DeathsPerTwhFromLowCarbonEnergySovacoolEtAl2016(ctx context.Context, entity string, year int) (*model.DeathsPerTwhFromLowCarbonEnergySovacoolEtAl2016Dataset, error) {
	var response model.DeathsPerTwhFromLowCarbonEnergySovacoolEtAl2016Dataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) DecompositionOfGenderWageGap1980BlauAndKahn2017(ctx context.Context, entity string, year int) (*model.DecompositionOfGenderWageGap1980BlauAndKahn2017Dataset, error) {
	var response model.DecompositionOfGenderWageGap1980BlauAndKahn2017Dataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) DecompositionOfGenderWageGap2010BlauAndKahn2017(ctx context.Context, entity string, year int) (*model.DecompositionOfGenderWageGap2010BlauAndKahn2017Dataset, error) {
	var response model.DecompositionOfGenderWageGap2010BlauAndKahn2017Dataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) DecompositionTimesOfMarineDebris(ctx context.Context, entity string, year int) (*model.DecompositionTimesOfMarineDebrisDataset, error) {
	var response model.DecompositionTimesOfMarineDebrisDataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) DeforestationByCommodityPendrillEtAl2019(ctx context.Context, entity string, year int) (*model.DeforestationByCommodityPendrillEtAl2019Dataset, error) {
	var response model.DeforestationByCommodityPendrillEtAl2019Dataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) DeforestationInTradePendrill(ctx context.Context, entity string, year int) (*model.DeforestationInTradePendrillDataset, error) {
	var response model.DeforestationInTradePendrillDataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) DeliveryPointsInTheUsUnitedStatesPostalService2018(ctx context.Context, entity string, year int) (*model.DeliveryPointsInTheUsUnitedStatesPostalService2018Dataset, error) {
	var response model.DeliveryPointsInTheUsUnitedStatesPostalService2018Dataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) DepressionPrevalenceByEducationOecd(ctx context.Context, entity string, year int) (*model.DepressionPrevalenceByEducationOecdDataset, error) {
	var response model.DepressionPrevalenceByEducationOecdDataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) DietCompositionsByCommodityCategoriesFao2017(ctx context.Context, entity string, year int) (*model.DietCompositionsByCommodityCategoriesFao2017Dataset, error) {
	var response model.DietCompositionsByCommodityCategoriesFao2017Dataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) DietCompositionsBySpecificFoodCommoditiesFao2017(ctx context.Context, entity string, year int) (*model.DietCompositionsBySpecificFoodCommoditiesFao2017Dataset, error) {
	var response model.DietCompositionsBySpecificFoodCommoditiesFao2017Dataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) DietaryMacronutrientCompositionsFao2017(ctx context.Context, entity string, year int) (*model.DietaryMacronutrientCompositionsFao2017Dataset, error) {
	var response model.DietaryMacronutrientCompositionsFao2017Dataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) DifferenceInTheValueOfGoodsExportedToAndImportedByTheUsFor2016Dots2017(ctx context.Context, entity string, year int) (*model.DifferenceInTheValueOfGoodsExportedToAndImportedByTheUsFor2016Dots2017Dataset, error) {
	var response model.DifferenceInTheValueOfGoodsExportedToAndImportedByTheUsFor2016Dots2017Dataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) DifferencesInPopulationEstimatesOwidBasedOnUnVsUsCensusBureau(ctx context.Context, entity string, year int) (*model.DifferencesInPopulationEstimatesOwidBasedOnUnVsUsCensusBureauDataset, error) {
	var response model.DifferencesInPopulationEstimatesOwidBasedOnUnVsUsCensusBureauDataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) DisabilityAdjustedLifeYearsWho2015(ctx context.Context, entity string, year int) (*model.DisabilityAdjustedLifeYearsWho2015Dataset, error) {
	var response model.DisabilityAdjustedLifeYearsWho2015Dataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) DistributionOfBilateralAndUnilateralTradePartnershipsFouquinAndHugotCepii2016(ctx context.Context, entity string, year int) (*model.DistributionOfBilateralAndUnilateralTradePartnershipsFouquinAndHugotCepii2016Dataset, error) {
	var response model.DistributionOfBilateralAndUnilateralTradePartnershipsFouquinAndHugotCepii2016Dataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) DrinkingHabitsInGreatBritainUkOns(ctx context.Context, entity string, year int) (*model.DrinkingHabitsInGreatBritainUkOnsDataset, error) {
	var response model.DrinkingHabitsInGreatBritainUkOnsDataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) DriversOfForestLossInBrazilLegalAmazonTyukavinaEtAl2017(ctx context.Context, entity string, year int) (*model.DriversOfForestLossInBrazilLegalAmazonTyukavinaEtAl2017Dataset, error) {
	var response model.DriversOfForestLossInBrazilLegalAmazonTyukavinaEtAl2017Dataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) DroughtSeverityIndexInUsNoaa(ctx context.Context, entity string, year int) (*model.DroughtSeverityIndexInUsNoaaDataset, error) {
	var response model.DroughtSeverityIndexInUsNoaaDataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) DurationOfMarriagesEndingInDivorceOwidBasedOnNationalStatistics(ctx context.Context, entity string, year int) (*model.DurationOfMarriagesEndingInDivorceOwidBasedOnNationalStatisticsDataset, error) {
	var response model.DurationOfMarriagesEndingInDivorceOwidBasedOnNationalStatisticsDataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) EciCountryRankingsObservatoryOfEconomicComplexity2016AndTheAtlasOfEconomicComplexity2016(ctx context.Context, entity string, year int) (*model.EciCountryRankingsObservatoryOfEconomicComplexity2016AndTheAtlasOfEconomicComplexity2016Dataset, error) {
	var response model.EciCountryRankingsObservatoryOfEconomicComplexity2016AndTheAtlasOfEconomicComplexity2016Dataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) EarthquakeDeathsNgdcNoaa(ctx context.Context, entity string, year int) (*model.EarthquakeDeathsNgdcNoaaDataset, error) {
	var response model.EarthquakeDeathsNgdcNoaaDataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) EconomicFreedomOfTheWorldFraserInstitute2018(ctx context.Context, entity string, year int) (*model.EconomicFreedomOfTheWorldFraserInstitute2018Dataset, error) {
	var response model.EconomicFreedomOfTheWorldFraserInstitute2018Dataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) EconomicImpacts2vs15cPretisEtAl(ctx context.Context, entity string, year int) (*model.EconomicImpacts2vs15cPretisEtAlDataset, error) {
	var response model.EconomicImpacts2vs15cPretisEtAlDataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) EconomicImpactsOf15cPretisEtAl(ctx context.Context, entity string, year int) (*model.EconomicImpactsOf15cPretisEtAlDataset, error) {
	var response model.EconomicImpactsOf15cPretisEtAlDataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) EconomicImpactsOf2cPretisEtAl2018(ctx context.Context, entity string, year int) (*model.EconomicImpactsOf2cPretisEtAl2018Dataset, error) {
	var response model.EconomicImpactsOf2cPretisEtAl2018Dataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) EconomicLossesFromDisastersAsAShareOfGdpPielke2018(ctx context.Context, entity string, year int) (*model.EconomicLossesFromDisastersAsAShareOfGdpPielke2018Dataset, error) {
	var response model.EconomicLossesFromDisastersAsAShareOfGdpPielke2018Dataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) EducationExpenditureUs1949_2013Nces2014(ctx context.Context, entity string, year int) (*model.EducationExpenditureUs19492013Nces2014Dataset, error) {
	var response model.EducationExpenditureUs19492013Nces2014Dataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) EducationDataDeprivationGemReport201718Uis2017(ctx context.Context, entity string, year int) (*model.EducationDataDeprivationGemReport201718Uis2017Dataset, error) {
	var response model.EducationDataDeprivationGemReport201718Uis2017Dataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) EducationalAttainmentBarroLeeEducationDataset2010(ctx context.Context, entity string, year int) (*model.EducationalAttainmentBarroLeeEducationDataset2010Dataset, error) {
	var response model.EducationalAttainmentBarroLeeEducationDataset2010Dataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) EducationalOutcomesHanushekAndWoessmann2012(ctx context.Context, entity string, year int) (*model.EducationalOutcomesHanushekAndWoessmann2012Dataset, error) {
	var response model.EducationalOutcomesHanushekAndWoessmann2012Dataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) ElectricityMixFromBpAndEmber2022(ctx context.Context, entity string, year int) (*model.ElectricityMixFromBpAndEmber2022Dataset, error) {
	var response model.ElectricityMixFromBpAndEmber2022Dataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) ElectricityMixFromBpAndEmber2022Archive(ctx context.Context, entity string, year int) (*model.ElectricityMixFromBpAndEmber2022ArchiveDataset, error) {
	var response model.ElectricityMixFromBpAndEmber2022ArchiveDataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) ElephantPopulationAfesgAndasesg2019(ctx context.Context, entity string, year int) (*model.ElephantPopulationAfesgAndasesg2019Dataset, error) {
	var response model.ElephantPopulationAfesgAndasesg2019Dataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) EmissionsAirPollutantsOverLongTermDefraAndEpa(ctx context.Context, entity string, year int) (*model.EmissionsAirPollutantsOverLongTermDefraAndEpaDataset, error) {
	var response model.EmissionsAirPollutantsOverLongTermDefraAndEpaDataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) EmissionsIntensityAndValueAddedBySectorLinusEtAl(ctx context.Context, entity string, year int) (*model.EmissionsIntensityAndValueAddedBySectorLinusEtAlDataset, error) {
	var response model.EmissionsIntensityAndValueAddedBySectorLinusEtAlDataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) EmploymentAndGenderAttitudesPewResearchCentre2012(ctx context.Context, entity string, year int) (*model.EmploymentAndGenderAttitudesPewResearchCentre2012Dataset, error) {
	var response model.EmploymentAndGenderAttitudesPewResearchCentre2012Dataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) EmploymentRateAges25_34ByEducationEducationAtAGlanceOecdIndicators2017(ctx context.Context, entity string, year int) (*model.EmploymentRateAges2534ByEducationEducationAtAGlanceOecdIndicators2017Dataset, error) {
	var response model.EmploymentRateAges2534ByEducationEducationAtAGlanceOecdIndicators2017Dataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) EndemicAndThreatenedInvertebrateSpeciesByCountryIucn2020(ctx context.Context, entity string, year int) (*model.EndemicAndThreatenedInvertebrateSpeciesByCountryIucn2020Dataset, error) {
	var response model.EndemicAndThreatenedInvertebrateSpeciesByCountryIucn2020Dataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) EndemicVertebrateSpeciesByCountryIucn2020(ctx context.Context, entity string, year int) (*model.EndemicVertebrateSpeciesByCountryIucn2020Dataset, error) {
	var response model.EndemicVertebrateSpeciesByCountryIucn2020Dataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) EnergyPricesBpStatistics2016(ctx context.Context, entity string, year int) (*model.EnergyPricesBpStatistics2016Dataset, error) {
	var response model.EnergyPricesBpStatistics2016Dataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) EnergyEfficiencyByPassengerModeInUsaBts(ctx context.Context, entity string, year int) (*model.EnergyEfficiencyByPassengerModeInUsaBtsDataset, error) {
	var response model.EnergyEfficiencyByPassengerModeInUsaBtsDataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) EnergyImportsPercEnergyUseWorldBank2014(ctx context.Context, entity string, year int) (*model.EnergyImportsPercEnergyUseWorldBank2014Dataset, error) {
	var response model.EnergyImportsPercEnergyUseWorldBank2014Dataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) EnergyLandUseScenarioAnalysisOwidBasedOnUneceAndEmber(ctx context.Context, entity string, year int) (*model.EnergyLandUseScenarioAnalysisOwidBasedOnUneceAndEmberDataset, error) {
	var response model.EnergyLandUseScenarioAnalysisOwidBasedOnUneceAndEmberDataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) EnergyMixFromBp2020(ctx context.Context, entity string, year int) (*model.EnergyMixFromBp2020Dataset, error) {
	var response model.EnergyMixFromBp2020Dataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) EnergyMixFromBp2021(ctx context.Context, entity string, year int) (*model.EnergyMixFromBp2021Dataset, error) {
	var response model.EnergyMixFromBp2021Dataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) EnergyMixInTheUkDukes2018(ctx context.Context, entity string, year int) (*model.EnergyMixInTheUkDukes2018Dataset, error) {
	var response model.EnergyMixInTheUkDukes2018Dataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) EnvironmentalImpactsOfFoodPooreAndNemecek2018(ctx context.Context, entity string, year int) (*model.EnvironmentalImpactsOfFoodPooreAndNemecek2018Dataset, error) {
	var response model.EnvironmentalImpactsOfFoodPooreAndNemecek2018Dataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) EstimatedAverageAgeAtMarriageByGenderUnAndOecd(ctx context.Context, entity string, year int) (*model.EstimatedAverageAgeAtMarriageByGenderUnAndOecdDataset, error) {
	var response model.EstimatedAverageAgeAtMarriageByGenderUnAndOecdDataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) EstimatedAverageAgeAtMarriageByGenderUn(ctx context.Context, entity string, year int) (*model.EstimatedAverageAgeAtMarriageByGenderUnDataset, error) {
	var response model.EstimatedAverageAgeAtMarriageByGenderUnDataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) EstimatedFundingAndFutureNeedsForHivInLowAndMiddleIncomeCountriesUnaids(ctx context.Context, entity string, year int) (*model.EstimatedFundingAndFutureNeedsForHivInLowAndMiddleIncomeCountriesUnaidsDataset, error) {
	var response model.EstimatedFundingAndFutureNeedsForHivInLowAndMiddleIncomeCountriesUnaidsDataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) EstimatedHistoricalLiteracyRatesBuringhAndVanZanden2009(ctx context.Context, entity string, year int) (*model.EstimatedHistoricalLiteracyRatesBuringhAndVanZanden2009Dataset, error) {
	var response model.EstimatedHistoricalLiteracyRatesBuringhAndVanZanden2009Dataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) EstimatedPercentOfWomenWhoAreMarriedOrInAUnionUn(ctx context.Context, entity string, year int) (*model.EstimatedPercentOfWomenWhoAreMarriedOrInAUnionUnDataset, error) {
	var response model.EstimatedPercentOfWomenWhoAreMarriedOrInAUnionUnDataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) EthnographicAndArchaeologicalEvidenceOnViolentDeaths(ctx context.Context, entity string, year int) (*model.EthnographicAndArchaeologicalEvidenceOnViolentDeathsDataset, error) {
	var response model.EthnographicAndArchaeologicalEvidenceOnViolentDeathsDataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) EuropeanVehiclePassengerSalesIcct2021(ctx context.Context, entity string, year int) (*model.EuropeanVehiclePassengerSalesIcct2021Dataset, error) {
	var response model.EuropeanVehiclePassengerSalesIcct2021Dataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) ExcessMortalityDataOwid2022(ctx context.Context, entity string, year int) (*model.ExcessMortalityDataOwid2022Dataset, error) {
	var response model.ExcessMortalityDataOwid2022Dataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) ExecutionsByCountryAmnestyInternational(ctx context.Context, entity string, year int) (*model.ExecutionsByCountryAmnestyInternationalDataset, error) {
	var response model.ExecutionsByCountryAmnestyInternationalDataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) ExpectedYearsOfSchoolingUndp2018(ctx context.Context, entity string, year int) (*model.ExpectedYearsOfSchoolingUndp2018Dataset, error) {
	var response model.ExpectedYearsOfSchoolingUndp2018Dataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) ExperienceCurvesLafond2017(ctx context.Context, entity string, year int) (*model.ExperienceCurvesLafond2017Dataset, error) {
	var response model.ExperienceCurvesLafond2017Dataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) ExtensionsInLifeExpectancyOwidCalculationsBasedOnUnPopulationDivision2017Revision(ctx context.Context, entity string, year int) (*model.ExtensionsInLifeExpectancyOwidCalculationsBasedOnUnPopulationDivision2017RevisionDataset, error) {
	var response model.ExtensionsInLifeExpectancyOwidCalculationsBasedOnUnPopulationDivision2017RevisionDataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) ExtremePoverty2030ProjectionsBySspCrespoEtAl2018(ctx context.Context, entity string, year int) (*model.ExtremePoverty2030ProjectionsBySspCrespoEtAl2018Dataset, error) {
	var response model.ExtremePoverty2030ProjectionsBySspCrespoEtAl2018Dataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) ExtremePovertyMichailMoatsosOecd(ctx context.Context, entity string, year int) (*model.ExtremePovertyMichailMoatsosOecdDataset, error) {
	var response model.ExtremePovertyMichailMoatsosOecdDataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) ExtremePovertyInAbsoluteNumbersRavallion2016UpdatedWithWorldBank2019(ctx context.Context, entity string, year int) (*model.ExtremePovertyInAbsoluteNumbersRavallion2016UpdatedWithWorldBank2019Dataset, error) {
	var response model.ExtremePovertyInAbsoluteNumbersRavallion2016UpdatedWithWorldBank2019Dataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) ExtremeTemperaturesInUsNoaa(ctx context.Context, entity string, year int) (*model.ExtremeTemperaturesInUsNoaaDataset, error) {
	var response model.ExtremeTemperaturesInUsNoaaDataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) ExtremeIncomePovertyInEuropeBradshawAndMayhew2011(ctx context.Context, entity string, year int) (*model.ExtremeIncomePovertyInEuropeBradshawAndMayhew2011Dataset, error) {
	var response model.ExtremeIncomePovertyInEuropeBradshawAndMayhew2011Dataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) ExtremePrecipitationInUsNoaa(ctx context.Context, entity string, year int) (*model.ExtremePrecipitationInUsNoaaDataset, error) {
	var response model.ExtremePrecipitationInUsNoaaDataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) Fao2030_50ProjectionsOfArableLandFao2017(ctx context.Context, entity string, year int) (*model.Fao203050ProjectionsOfArableLandFao2017Dataset, error) {
	var response model.Fao203050ProjectionsOfArableLandFao2017Dataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) FaoUndernourishmentComparison2010Vs2012(ctx context.Context, entity string, year int) (*model.FaoUndernourishmentComparison2010Vs2012Dataset, error) {
	var response model.FaoUndernourishmentComparison2010Vs2012Dataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) FamilyBenefitsPublicSpendingOecd2016(ctx context.Context, entity string, year int) (*model.FamilyBenefitsPublicSpendingOecd2016Dataset, error) {
	var response model.FamilyBenefitsPublicSpendingOecd2016Dataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) FatalAviationAccidentsAndFataltiesAviationSafetyNetworkAsn(ctx context.Context, entity string, year int) (*model.FatalAviationAccidentsAndFataltiesAviationSafetyNetworkAsnDataset, error) {
	var response model.FatalAviationAccidentsAndFataltiesAviationSafetyNetworkAsnDataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) FemaleAndMaleLifeExpectancyAtBirthOwidBasedOnUnPopulationDivision2017(ctx context.Context, entity string, year int) (*model.FemaleAndMaleLifeExpectancyAtBirthOwidBasedOnUnPopulationDivision2017Dataset, error) {
	var response model.FemaleAndMaleLifeExpectancyAtBirthOwidBasedOnUnPopulationDivision2017Dataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) FemaleLaborForceParticipationRateOwid2017(ctx context.Context, entity string, year int) (*model.FemaleLaborForceParticipationRateOwid2017Dataset, error) {
	var response model.FemaleLaborForceParticipationRateOwid2017Dataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) FemaleWeeklyHoursWorkedOecd2017(ctx context.Context, entity string, year int) (*model.FemaleWeeklyHoursWorkedOecd2017Dataset, error) {
	var response model.FemaleWeeklyHoursWorkedOecd2017Dataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) FertilityUnPopulationDivision2015Revision(ctx context.Context, entity string, year int) (*model.FertilityUnPopulationDivision2015RevisionDataset, error) {
	var response model.FertilityUnPopulationDivision2015RevisionDataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) FertilityRateSelectedGapminderV12_2017(ctx context.Context, entity string, year int) (*model.FertilityRateSelectedGapminderV122017Dataset, error) {
	var response model.FertilityRateSelectedGapminderV122017Dataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) FertilityRateWcIiasa2017(ctx context.Context, entity string, year int) (*model.FertilityRateWcIiasa2017Dataset, error) {
	var response model.FertilityRateWcIiasa2017Dataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) FertilityRateCompleteGapminderV12_2017(ctx context.Context, entity string, year int) (*model.FertilityRateCompleteGapminderV122017Dataset, error) {
	var response model.FertilityRateCompleteGapminderV122017Dataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) FertilizerPricesWorldBank2017(ctx context.Context, entity string, year int) (*model.FertilizerPricesWorldBank2017Dataset, error) {
	var response model.FertilizerPricesWorldBank2017Dataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) FertilizerUsePerHectareOfLandFaoAndFederico(ctx context.Context, entity string, year int) (*model.FertilizerUsePerHectareOfLandFaoAndFedericoDataset, error) {
	var response model.FertilizerUsePerHectareOfLandFaoAndFedericoDataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) FirmsWithFinancialConstraintsWorldBankEnterpriseSurvey2019(ctx context.Context, entity string, year int) (*model.FirmsWithFinancialConstraintsWorldBankEnterpriseSurvey2019Dataset, error) {
	var response model.FirmsWithFinancialConstraintsWorldBankEnterpriseSurvey2019Dataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) FishStocksRamlegacy(ctx context.Context, entity string, year int) (*model.FishStocksRamlegacyDataset, error) {
	var response model.FishStocksRamlegacyDataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) FisheryCatchBreakdownPaulyAndZeller2016(ctx context.Context, entity string, year int) (*model.FisheryCatchBreakdownPaulyAndZeller2016Dataset, error) {
	var response model.FisheryCatchBreakdownPaulyAndZeller2016Dataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) FiveYearCancerSurvivalRatesAllemaniEtAl(ctx context.Context, entity string, year int) (*model.FiveYearCancerSurvivalRatesAllemaniEtAlDataset, error) {
	var response model.FiveYearCancerSurvivalRatesAllemaniEtAlDataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) FiveYearCancerSurvivalRatesNationalCancerInstitute(ctx context.Context, entity string, year int) (*model.FiveYearCancerSurvivalRatesNationalCancerInstituteDataset, error) {
	var response model.FiveYearCancerSurvivalRatesNationalCancerInstituteDataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) FoodPricesExpressedInHourlyWagesUsBureauOfLaborStatistics2015(ctx context.Context, entity string, year int) (*model.FoodPricesExpressedInHourlyWagesUsBureauOfLaborStatistics2015Dataset, error) {
	var response model.FoodPricesExpressedInHourlyWagesUsBureauOfLaborStatistics2015Dataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) FoodExpenditureInTheUsaUsda(ctx context.Context, entity string, year int) (*model.FoodExpenditureInTheUsaUsdaDataset, error) {
	var response model.FoodExpenditureInTheUsaUsdaDataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) FoodMilesByTransportMethodPooreAndNemecek2018(ctx context.Context, entity string, year int) (*model.FoodMilesByTransportMethodPooreAndNemecek2018Dataset, error) {
	var response model.FoodMilesByTransportMethodPooreAndNemecek2018Dataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) FoodSupplyFao2020(ctx context.Context, entity string, year int) (*model.FoodSupplyFao2020Dataset, error) {
	var response model.FoodSupplyFao2020Dataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) FoodWasteInTheEuropeanUnionEuropa2010(ctx context.Context, entity string, year int) (*model.FoodWasteInTheEuropeanUnionEuropa2010Dataset, error) {
	var response model.FoodWasteInTheEuropeanUnionEuropa2010Dataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) FoodWasteInTheSupplyChainTWrap2015AndEuropa2015(ctx context.Context, entity string, year int) (*model.FoodWasteInTheSupplyChainTWrap2015AndEuropa2015Dataset, error) {
	var response model.FoodWasteInTheSupplyChainTWrap2015AndEuropa2015Dataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) ForestTransitionPhasePendrillEtAl2019(ctx context.Context, entity string, year int) (*model.ForestTransitionPhasePendrillEtAl2019Dataset, error) {
	var response model.ForestTransitionPhasePendrillEtAl2019Dataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) ForestLandDeforestationAndChangeFao2020(ctx context.Context, entity string, year int) (*model.ForestLandDeforestationAndChangeFao2020Dataset, error) {
	var response model.ForestLandDeforestationAndChangeFao2020Dataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) ForestryAreaFao2017(ctx context.Context, entity string, year int) (*model.ForestryAreaFao2017Dataset, error) {
	var response model.ForestryAreaFao2017Dataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) FossilFuelConsumptionPerCapitaBpAndUn2017Revision(ctx context.Context, entity string, year int) (*model.FossilFuelConsumptionPerCapitaBpAndUn2017RevisionDataset, error) {
	var response model.FossilFuelConsumptionPerCapitaBpAndUn2017RevisionDataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) FossilFuelProductionBpAndShift2020(ctx context.Context, entity string, year int) (*model.FossilFuelProductionBpAndShift2020Dataset, error) {
	var response model.FossilFuelProductionBpAndShift2020Dataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) FossilFuelProductionBpAndShift2022(ctx context.Context, entity string, year int) (*model.FossilFuelProductionBpAndShift2022Dataset, error) {
	var response model.FossilFuelProductionBpAndShift2022Dataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) FriendsAndFamilySupportOecdBasedOnGallup2016(ctx context.Context, entity string, year int) (*model.FriendsAndFamilySupportOecdBasedOnGallup2016Dataset, error) {
	var response model.FriendsAndFamilySupportOecdBasedOnGallup2016Dataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) GdpGrowthFromPreviousYear2020Q2EurostatOecdNationalSources(ctx context.Context, entity string, year int) (*model.GdpGrowthFromPreviousYear2020Q2EurostatOecdNationalSourcesDataset, error) {
	var response model.GdpGrowthFromPreviousYear2020Q2EurostatOecdNationalSourcesDataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) GdpInEnglandUsingBoe2017(ctx context.Context, entity string, year int) (*model.GdpInEnglandUsingBoe2017Dataset, error) {
	var response model.GdpInEnglandUsingBoe2017Dataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) GdpPerCapitaIndexedAt1950MaddisonProjectData2018(ctx context.Context, entity string, year int) (*model.GdpPerCapitaIndexedAt1950MaddisonProjectData2018Dataset, error) {
	var response model.GdpPerCapitaIndexedAt1950MaddisonProjectData2018Dataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) GhgEmissionsByCountryAndSectorCait2020(ctx context.Context, entity string, year int) (*model.GhgEmissionsByCountryAndSectorCait2020Dataset, error) {
	var response model.GhgEmissionsByCountryAndSectorCait2020Dataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) GhgEmissionsByCountryAndSectorCait2021(ctx context.Context, entity string, year int) (*model.GhgEmissionsByCountryAndSectorCait2021Dataset, error) {
	var response model.GhgEmissionsByCountryAndSectorCait2021Dataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) GhgEmissionsPerCapitaEdgar2019(ctx context.Context, entity string, year int) (*model.GhgEmissionsPerCapitaEdgar2019Dataset, error) {
	var response model.GhgEmissionsPerCapitaEdgar2019Dataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) GistempTemperatureAnomaly(ctx context.Context, entity string, year int) (*model.GistempTemperatureAnomalyDataset, error) {
	var response model.GistempTemperatureAnomalyDataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) GapminderIgnoranceTestResultsGapminder(ctx context.Context, entity string, year int) (*model.GapminderIgnoranceTestResultsGapminderDataset, error) {
	var response model.GapminderIgnoranceTestResultsGapminderDataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) GasProductionEtemadAndLuciana(ctx context.Context, entity string, year int) (*model.GasProductionEtemadAndLucianaDataset, error) {
	var response model.GasProductionEtemadAndLucianaDataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) GenderInequalityIndexHumanDevelopmentReport2015(ctx context.Context, entity string, year int) (*model.GenderInequalityIndexHumanDevelopmentReport2015Dataset, error) {
	var response model.GenderInequalityIndexHumanDevelopmentReport2015Dataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) GenderWageGapOecd2017(ctx context.Context, entity string, year int) (*model.GenderWageGapOecd2017Dataset, error) {
	var response model.GenderWageGapOecd2017Dataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) GenderPreferenceForBossGallup2017(ctx context.Context, entity string, year int) (*model.GenderPreferenceForBossGallup2017Dataset, error) {
	var response model.GenderPreferenceForBossGallup2017Dataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) GenderWageGapAssigningZerosForNoWork(ctx context.Context, entity string, year int) (*model.GenderWageGapAssigningZerosForNoWorkDataset, error) {
	var response model.GenderWageGapAssigningZerosForNoWorkDataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) GenuineSavingEstimatesByVariousMeasuresBlumDucoingMclaughlin2017(ctx context.Context, entity string, year int) (*model.GenuineSavingEstimatesByVariousMeasuresBlumDucoingMclaughlin2017Dataset, error) {
	var response model.GenuineSavingEstimatesByVariousMeasuresBlumDucoingMclaughlin2017Dataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) GermanRoadDeathsAndAccidentsDestatis(ctx context.Context, entity string, year int) (*model.GermanRoadDeathsAndAccidentsDestatisDataset, error) {
	var response model.GermanRoadDeathsAndAccidentsDestatisDataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) GiniCoefficientEquivalizedIncomeAfterTaxAndTransfersChartbookOfEconomicInequality2017(ctx context.Context, entity string, year int) (*model.GiniCoefficientEquivalizedIncomeAfterTaxAndTransfersChartbookOfEconomicInequality2017Dataset, error) {
	var response model.GiniCoefficientEquivalizedIncomeAfterTaxAndTransfersChartbookOfEconomicInequality2017Dataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) GiniCoefficientsOecd2016(ctx context.Context, entity string, year int) (*model.GiniCoefficientsOecd2016Dataset, error) {
	var response model.GiniCoefficientsOecd2016Dataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) GiniCoefficientsForLifetimeInequalityPeltzman2009(ctx context.Context, entity string, year int) (*model.GiniCoefficientsForLifetimeInequalityPeltzman2009Dataset, error) {
	var response model.GiniCoefficientsForLifetimeInequalityPeltzman2009Dataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) GlobalBmiInFemalesNcdrisc2017(ctx context.Context, entity string, year int) (*model.GlobalBmiInFemalesNcdrisc2017Dataset, error) {
	var response model.GlobalBmiInFemalesNcdrisc2017Dataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) GlobalBmiInMalesNcdRisc2017(ctx context.Context, entity string, year int) (*model.GlobalBmiInMalesNcdRisc2017Dataset, error) {
	var response model.GlobalBmiInMalesNcdRisc2017Dataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) GlobalCo2EmissionsCdiacAndUnPopulation(ctx context.Context, entity string, year int) (*model.GlobalCo2EmissionsCdiacAndUnPopulationDataset, error) {
	var response model.GlobalCo2EmissionsCdiacAndUnPopulationDataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) GlobalCarbonBudgetGcp2021(ctx context.Context, entity string, year int) (*model.GlobalCarbonBudgetGcp2021Dataset, error) {
	var response model.GlobalCarbonBudgetGcp2021Dataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) GlobalCarbonBudgetFor2cIpcc2013(ctx context.Context, entity string, year int) (*model.GlobalCarbonBudgetFor2cIpcc2013Dataset, error) {
	var response model.GlobalCarbonBudgetFor2cIpcc2013Dataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) GlobalDataSetOnEducationQuality1965_2015AltinokAngristAndPatrinos2018(ctx context.Context, entity string, year int) (*model.GlobalDataSetOnEducationQuality19652015AltinokAngristAndPatrinos2018Dataset, error) {
	var response model.GlobalDataSetOnEducationQuality19652015AltinokAngristAndPatrinos2018Dataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) GlobalHungerIndex2021(ctx context.Context, entity string, year int) (*model.GlobalHungerIndex2021Dataset, error) {
	var response model.GlobalHungerIndex2021Dataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) GlobalHungerIndexIn1992Listed2017GlobalHungerIndex2017(ctx context.Context, entity string, year int) (*model.GlobalHungerIndexIn1992Listed2017GlobalHungerIndex2017Dataset, error) {
	var response model.GlobalHungerIndexIn1992Listed2017GlobalHungerIndex2017Dataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) GlobalHungerIndexIn2017Listed2017GlobalHungerIndex2017(ctx context.Context, entity string, year int) (*model.GlobalHungerIndexIn2017Listed2017GlobalHungerIndex2017Dataset, error) {
	var response model.GlobalHungerIndexIn2017Listed2017GlobalHungerIndex2017Dataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) GlobalHungerIndexIfpri2018(ctx context.Context, entity string, year int) (*model.GlobalHungerIndexIfpri2018Dataset, error) {
	var response model.GlobalHungerIndexIfpri2018Dataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) GlobalPopulationByRegionWithProjectionsHyde2016AndUn2017(ctx context.Context, entity string, year int) (*model.GlobalPopulationByRegionWithProjectionsHyde2016AndUn2017Dataset, error) {
	var response model.GlobalPopulationByRegionWithProjectionsHyde2016AndUn2017Dataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) GlobalPrimaryEnergyConsumptionVaclavSmil2017AndBpStatistics2020(ctx context.Context, entity string, year int) (*model.GlobalPrimaryEnergyConsumptionVaclavSmil2017AndBpStatistics2020Dataset, error) {
	var response model.GlobalPrimaryEnergyConsumptionVaclavSmil2017AndBpStatistics2020Dataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) GlobalProjectionMediumSsp2Iiasa2016(ctx context.Context, entity string, year int) (*model.GlobalProjectionMediumSsp2Iiasa2016Dataset, error) {
	var response model.GlobalProjectionMediumSsp2Iiasa2016Dataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) GlobalRevenueStatisticsDatabaseOecd2018(ctx context.Context, entity string, year int) (*model.GlobalRevenueStatisticsDatabaseOecd2018Dataset, error) {
	var response model.GlobalRevenueStatisticsDatabaseOecd2018Dataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) GlobalSmallpoxCases(ctx context.Context, entity string, year int) (*model.GlobalSmallpoxCasesDataset, error) {
	var response model.GlobalSmallpoxCasesDataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) GlobalTuberculosisReportCaseNotificationsWho2019(ctx context.Context, entity string, year int) (*model.GlobalTuberculosisReportCaseNotificationsWho2019Dataset, error) {
	var response model.GlobalTuberculosisReportCaseNotificationsWho2019Dataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) GlobalTuberculosisReportTbBurdenEstimatesWho2019(ctx context.Context, entity string, year int) (*model.GlobalTuberculosisReportTbBurdenEstimatesWho2019Dataset, error) {
	var response model.GlobalTuberculosisReportTbBurdenEstimatesWho2019Dataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) GlobalAgriculturalLandByCropFao2017(ctx context.Context, entity string, year int) (*model.GlobalAgriculturalLandByCropFao2017Dataset, error) {
	var response model.GlobalAgriculturalLandByCropFao2017Dataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) GlobalAirlineTrafficAndCapacityIcao2020(ctx context.Context, entity string, year int) (*model.GlobalAirlineTrafficAndCapacityIcao2020Dataset, error) {
	var response model.GlobalAirlineTrafficAndCapacityIcao2020Dataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) GlobalAverageTemperatureAnomalyHadleyCentre(ctx context.Context, entity string, year int) (*model.GlobalAverageTemperatureAnomalyHadleyCentreDataset, error) {
	var response model.GlobalAverageTemperatureAnomalyHadleyCentreDataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) GlobalChildMortalitySince1800BasedOnGapminderAndWorldBank2019(ctx context.Context, entity string, year int) (*model.GlobalChildMortalitySince1800BasedOnGapminderAndWorldBank2019Dataset, error) {
	var response model.GlobalChildMortalitySince1800BasedOnGapminderAndWorldBank2019Dataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) GlobalDeathRatesFromDisastersEmdatUnAndHyde(ctx context.Context, entity string, year int) (*model.GlobalDeathRatesFromDisastersEmdatUnAndHydeDataset, error) {
	var response model.GlobalDeathRatesFromDisastersEmdatUnAndHydeDataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) GlobalDeathsByCauseAndRiskGlobalBurdenOfDisease2017(ctx context.Context, entity string, year int) (*model.GlobalDeathsByCauseAndRiskGlobalBurdenOfDisease2017Dataset, error) {
	var response model.GlobalDeathsByCauseAndRiskGlobalBurdenOfDisease2017Dataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) GlobalEducationOecdIiasa2016(ctx context.Context, entity string, year int) (*model.GlobalEducationOecdIiasa2016Dataset, error) {
	var response model.GlobalEducationOecdIiasa2016Dataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) GlobalFishCatchByEndUseFishstatViaSeaaroundus(ctx context.Context, entity string, year int) (*model.GlobalFishCatchByEndUseFishstatViaSeaaroundusDataset, error) {
	var response model.GlobalFishCatchByEndUseFishstatViaSeaaroundusDataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) GlobalFreshwaterUseSince1900Igb(ctx context.Context, entity string, year int) (*model.GlobalFreshwaterUseSince1900IgbDataset, error) {
	var response model.GlobalFreshwaterUseSince1900IgbDataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) GlobalLiteracySince1800OwidBasedOnOecdAndUnesco2019(ctx context.Context, entity string, year int) (*model.GlobalLiteracySince1800OwidBasedOnOecdAndUnesco2019Dataset, error) {
	var response model.GlobalLiteracySince1800OwidBasedOnOecdAndUnesco2019Dataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) GlobalMeatProjectionsTo2050Fao(ctx context.Context, entity string, year int) (*model.GlobalMeatProjectionsTo2050FaoDataset, error) {
	var response model.GlobalMeatProjectionsTo2050FaoDataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) GlobalPlasticProductionGeyerEtAl2017(ctx context.Context, entity string, year int) (*model.GlobalPlasticProductionGeyerEtAl2017Dataset, error) {
	var response model.GlobalPlasticProductionGeyerEtAl2017Dataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) GlobalPopulationTrendsUsCensusBureau2017(ctx context.Context, entity string, year int) (*model.GlobalPopulationTrendsUsCensusBureau2017Dataset, error) {
	var response model.GlobalPopulationTrendsUsCensusBureau2017Dataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) GlobalPrecipitationAnomalyNoaa(ctx context.Context, entity string, year int) (*model.GlobalPrecipitationAnomalyNoaaDataset, error) {
	var response model.GlobalPrecipitationAnomalyNoaaDataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) GlobalPrimaryEnergyShareSmilAndBp(ctx context.Context, entity string, year int) (*model.GlobalPrimaryEnergyShareSmilAndBpDataset, error) {
	var response model.GlobalPrimaryEnergyShareSmilAndBpDataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) GlobalTemperatureAnomalyMetOfficeHadcrut4(ctx context.Context, entity string, year int) (*model.GlobalTemperatureAnomalyMetOfficeHadcrut4Dataset, error) {
	var response model.GlobalTemperatureAnomalyMetOfficeHadcrut4Dataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) GlobalWarmingPotentialFactorsGwp100Ipcc2014(ctx context.Context, entity string, year int) (*model.GlobalWarmingPotentialFactorsGwp100Ipcc2014Dataset, error) {
	var response model.GlobalWarmingPotentialFactorsGwp100Ipcc2014Dataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) GlobalYearOfLastPolioCasePlusCertificationStatusGpei2017(ctx context.Context, entity string, year int) (*model.GlobalYearOfLastPolioCasePlusCertificationStatusGpei2017Dataset, error) {
	var response model.GlobalYearOfLastPolioCasePlusCertificationStatusGpei2017Dataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) GlobalizationOver5CenturiesPwt90KlasingAndMilionis2014AndEstevadeordalFrantzAndTaylor2003(ctx context.Context, entity string, year int) (*model.GlobalizationOver5CenturiesPwt90KlasingAndMilionis2014AndEstevadeordalFrantzAndTaylor2003Dataset, error) {
	var response model.GlobalizationOver5CenturiesPwt90KlasingAndMilionis2014AndEstevadeordalFrantzAndTaylor2003Dataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) GoldPricesLaurenceAndWilliamson2017(ctx context.Context, entity string, year int) (*model.GoldPricesLaurenceAndWilliamson2017Dataset, error) {
	var response model.GoldPricesLaurenceAndWilliamson2017Dataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) GoogleMobilityTrends2020(ctx context.Context, entity string, year int) (*model.GoogleMobilityTrends2020Dataset, error) {
	var response model.GoogleMobilityTrends2020Dataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) GovernmentEducationExpenditure1960_2010Szirmai2015(ctx context.Context, entity string, year int) (*model.GovernmentEducationExpenditure19602010Szirmai2015Dataset, error) {
	var response model.GovernmentEducationExpenditure19602010Szirmai2015Dataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) GovernmentExpenditureImfBasedOnMauroEtAl2015(ctx context.Context, entity string, year int) (*model.GovernmentExpenditureImfBasedOnMauroEtAl2015Dataset, error) {
	var response model.GovernmentExpenditureImfBasedOnMauroEtAl2015Dataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) GovernmentRevenueWallis2000(ctx context.Context, entity string, year int) (*model.GovernmentRevenueWallis2000Dataset, error) {
	var response model.GovernmentRevenueWallis2000Dataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) GovernmentSpendingRoineVlachosWaldenstrom2009AndUsHistoricalTables2016(ctx context.Context, entity string, year int) (*model.GovernmentSpendingRoineVlachosWaldenstrom2009AndUsHistoricalTables2016Dataset, error) {
	var response model.GovernmentSpendingRoineVlachosWaldenstrom2009AndUsHistoricalTables2016Dataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) GovernmentTransparencyIndexHollyerEtAl2014(ctx context.Context, entity string, year int) (*model.GovernmentTransparencyIndexHollyerEtAl2014Dataset, error) {
	var response model.GovernmentTransparencyIndexHollyerEtAl2014Dataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) GovernmentExpenditureAndLearningOutcomes(ctx context.Context, entity string, year int) (*model.GovernmentExpenditureAndLearningOutcomesDataset, error) {
	var response model.GovernmentExpenditureAndLearningOutcomesDataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) GovernmentSpendingOecd2017(ctx context.Context, entity string, year int) (*model.GovernmentSpendingOecd2017Dataset, error) {
	var response model.GovernmentSpendingOecd2017Dataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) GuineaWormCasesTheCarterCenter2022(ctx context.Context, entity string, year int) (*model.GuineaWormCasesTheCarterCenter2022Dataset, error) {
	var response model.GuineaWormCasesTheCarterCenter2022Dataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) GuineaWormCasesWho2018(ctx context.Context, entity string, year int) (*model.GuineaWormCasesWho2018Dataset, error) {
	var response model.GuineaWormCasesWho2018Dataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) GuineaWormWhoCertificationStatusWho2018(ctx context.Context, entity string, year int) (*model.GuineaWormWhoCertificationStatusWho2018Dataset, error) {
	var response model.GuineaWormWhoCertificationStatusWho2018Dataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) HalfIndexLandUseAlexanderEtAl2016(ctx context.Context, entity string, year int) (*model.HalfIndexLandUseAlexanderEtAl2016Dataset, error) {
	var response model.HalfIndexLandUseAlexanderEtAl2016Dataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) HadcrutTemperatureAnomaly(ctx context.Context, entity string, year int) (*model.HadcrutTemperatureAnomalyDataset, error) {
	var response model.HadcrutTemperatureAnomalyDataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) HappinessPredictorsWorldHappinessReport2017(ctx context.Context, entity string, year int) (*model.HappinessPredictorsWorldHappinessReport2017Dataset, error) {
	var response model.HappinessPredictorsWorldHappinessReport2017Dataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) HealthCoverageIlo2014(ctx context.Context, entity string, year int) (*model.HealthCoverageIlo2014Dataset, error) {
	var response model.HealthCoverageIlo2014Dataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) HealthExpenditureUk1950_2012OfficeOfHealthEconomics2012(ctx context.Context, entity string, year int) (*model.HealthExpenditureUk19502012OfficeOfHealthEconomics2012Dataset, error) {
	var response model.HealthExpenditureUk19502012OfficeOfHealthEconomics2012Dataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) HealthExpenditureUs1929_2013PrivateUsCensusAndWdi2013(ctx context.Context, entity string, year int) (*model.HealthExpenditureUs19292013PrivateUsCensusAndWdi2013Dataset, error) {
	var response model.HealthExpenditureUs19292013PrivateUsCensusAndWdi2013Dataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) HealthExpenditureUs1929_2013PublicUsCensusAndWdi2013(ctx context.Context, entity string, year int) (*model.HealthExpenditureUs19292013PublicUsCensusAndWdi2013Dataset, error) {
	var response model.HealthExpenditureUs19292013PublicUsCensusAndWdi2013Dataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) HealthExpenditureAndFinancingOecdstat2017(ctx context.Context, entity string, year int) (*model.HealthExpenditureAndFinancingOecdstat2017Dataset, error) {
	var response model.HealthExpenditureAndFinancingOecdstat2017Dataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) HealthInsuranceCoverageUsUsCurrentPopulationSurvey2014(ctx context.Context, entity string, year int) (*model.HealthInsuranceCoverageUsUsCurrentPopulationSurvey2014Dataset, error) {
	var response model.HealthInsuranceCoverageUsUsCurrentPopulationSurvey2014Dataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) HealthExpenditurePerCapitaWorldBankWdi2018(ctx context.Context, entity string, year int) (*model.HealthExpenditurePerCapitaWorldBankWdi2018Dataset, error) {
	var response model.HealthExpenditurePerCapitaWorldBankWdi2018Dataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) HealthProviderAbsenceRatesChaudhuryHammerKremerMuralidharanAndRogers2006(ctx context.Context, entity string, year int) (*model.HealthProviderAbsenceRatesChaudhuryHammerKremerMuralidharanAndRogers2006Dataset, error) {
	var response model.HealthProviderAbsenceRatesChaudhuryHammerKremerMuralidharanAndRogers2006Dataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) HealthcareAccessAndQualityIndexIhme2017(ctx context.Context, entity string, year int) (*model.HealthcareAccessAndQualityIndexIhme2017Dataset, error) {
	var response model.HealthcareAccessAndQualityIndexIhme2017Dataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) HealthyLifeExpectancyIhme(ctx context.Context, entity string, year int) (*model.HealthyLifeExpectancyIhmeDataset, error) {
	var response model.HealthyLifeExpectancyIhmeDataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) HeightsOfEarlyEuropeansBasedOnHermanussen2003AndTheNcdRisc2017(ctx context.Context, entity string, year int) (*model.HeightsOfEarlyEuropeansBasedOnHermanussen2003AndTheNcdRisc2017Dataset, error) {
	var response model.HeightsOfEarlyEuropeansBasedOnHermanussen2003AndTheNcdRisc2017Dataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) HiddenHungerIndexInPreSchoolChildrenMuthayyaEtAl2013(ctx context.Context, entity string, year int) (*model.HiddenHungerIndexInPreSchoolChildrenMuthayyaEtAl2013Dataset, error) {
	var response model.HiddenHungerIndexInPreSchoolChildrenMuthayyaEtAl2013Dataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) HistoricalIndexOfHumanDevelopmentWithoutGdpPradosDeLaEscosura(ctx context.Context, entity string, year int) (*model.HistoricalIndexOfHumanDevelopmentWithoutGdpPradosDeLaEscosuraDataset, error) {
	var response model.HistoricalIndexOfHumanDevelopmentWithoutGdpPradosDeLaEscosuraDataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) HistoricalIndexOfHumanDevelopmentPradosDeLaEscosura(ctx context.Context, entity string, year int) (*model.HistoricalIndexOfHumanDevelopmentPradosDeLaEscosuraDataset, error) {
	var response model.HistoricalIndexOfHumanDevelopmentPradosDeLaEscosuraDataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) HistoricalUnPopulationProjections(ctx context.Context, entity string, year int) (*model.HistoricalUnPopulationProjectionsDataset, error) {
	var response model.HistoricalUnPopulationProjectionsDataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) HistoricalCostOfComputerMemoryAndStorageJohnCMccallum(ctx context.Context, entity string, year int) (*model.HistoricalCostOfComputerMemoryAndStorageJohnCMccallumDataset, error) {
	var response model.HistoricalCostOfComputerMemoryAndStorageJohnCMccallumDataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) HistoricalEmploymentAndOutputBySectorOwid2017(ctx context.Context, entity string, year int) (*model.HistoricalEmploymentAndOutputBySectorOwid2017Dataset, error) {
	var response model.HistoricalEmploymentAndOutputBySectorOwid2017Dataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) HistoricalGenderEqualityIndexHowWasLife2014(ctx context.Context, entity string, year int) (*model.HistoricalGenderEqualityIndexHowWasLife2014Dataset, error) {
	var response model.HistoricalGenderEqualityIndexHowWasLife2014Dataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) HistoricalUrbanFractionEstimatesAndTotalComputedUrbanAreasHyde31_2010(ctx context.Context, entity string, year int) (*model.HistoricalUrbanFractionEstimatesAndTotalComputedUrbanAreasHyde312010Dataset, error) {
	var response model.HistoricalUrbanFractionEstimatesAndTotalComputedUrbanAreasHyde312010Dataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) HistoricalWorldPopulationComparisonOfDifferentSources(ctx context.Context, entity string, year int) (*model.HistoricalWorldPopulationComparisonOfDifferentSourcesDataset, error) {
	var response model.HistoricalWorldPopulationComparisonOfDifferentSourcesDataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) HomelessnessAndPrecariousHousingOecd2016(ctx context.Context, entity string, year int) (*model.HomelessnessAndPrecariousHousingOecd2016Dataset, error) {
	var response model.HomelessnessAndPrecariousHousingOecd2016Dataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) HomelessnessPrevalenceToroEtAl2007(ctx context.Context, entity string, year int) (*model.HomelessnessPrevalenceToroEtAl2007Dataset, error) {
	var response model.HomelessnessPrevalenceToroEtAl2007Dataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) HomicideRatesInEuropeOverLongTermEisnerAndIhme(ctx context.Context, entity string, year int) (*model.HomicideRatesInEuropeOverLongTermEisnerAndIhmeDataset, error) {
	var response model.HomicideRatesInEuropeOverLongTermEisnerAndIhmeDataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) HomosexualityLawsOwidBasedOnKennyAndPatel2017(ctx context.Context, entity string, year int) (*model.HomosexualityLawsOwidBasedOnKennyAndPatel2017Dataset, error) {
	var response model.HomosexualityLawsOwidBasedOnKennyAndPatel2017Dataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) HomosexualityOpinionsWvs1981_2016(ctx context.Context, entity string, year int) (*model.HomosexualityOpinionsWvs19812016Dataset, error) {
	var response model.HomosexualityOpinionsWvs19812016Dataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) HomosexualityPublicOpinionPewResearch2013(ctx context.Context, entity string, year int) (*model.HomosexualityPublicOpinionPewResearch2013Dataset, error) {
	var response model.HomosexualityPublicOpinionPewResearch2013Dataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) HouseholdExpenditureOnHousingWaterElectricityGasAndOtherFuelsAsAShareOfGdpUn(ctx context.Context, entity string, year int) (*model.HouseholdExpenditureOnHousingWaterElectricityGasAndOtherFuelsAsAShareOfGdpUnDataset, error) {
	var response model.HouseholdExpenditureOnHousingWaterElectricityGasAndOtherFuelsAsAShareOfGdpUnDataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) HouseholdsUsingSolidFuelsForCookingUrbanVsRuralUn(ctx context.Context, entity string, year int) (*model.HouseholdsUsingSolidFuelsForCookingUrbanVsRuralUnDataset, error) {
	var response model.HouseholdsUsingSolidFuelsForCookingUrbanVsRuralUnDataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) HouseholdsActualAndImputedRentAsShareOfGdpOecd(ctx context.Context, entity string, year int) (*model.HouseholdsActualAndImputedRentAsShareOfGdpOecdDataset, error) {
	var response model.HouseholdsActualAndImputedRentAsShareOfGdpOecdDataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) HowEuropeansSpendTheirTimeEuropeanCommission2004(ctx context.Context, entity string, year int) (*model.HowEuropeansSpendTheirTimeEuropeanCommission2004Dataset, error) {
	var response model.HowEuropeansSpendTheirTimeEuropeanCommission2004Dataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) HubbertsPeakCavalloAndEia(ctx context.Context, entity string, year int) (*model.HubbertsPeakCavalloAndEiaDataset, error) {
	var response model.HubbertsPeakCavalloAndEiaDataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) HumanCapitalIndexWorldBank2018(ctx context.Context, entity string, year int) (*model.HumanCapitalIndexWorldBank2018Dataset, error) {
	var response model.HumanCapitalIndexWorldBank2018Dataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) HumanCapitalInLongRunLeeLee2016(ctx context.Context, entity string, year int) (*model.HumanCapitalInLongRunLeeLee2016Dataset, error) {
	var response model.HumanCapitalInLongRunLeeLee2016Dataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) HumanDevelopmentIndexUndp(ctx context.Context, entity string, year int) (*model.HumanDevelopmentIndexUndpDataset, error) {
	var response model.HumanDevelopmentIndexUndpDataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) HumanHeightNcdRisc2017(ctx context.Context, entity string, year int) (*model.HumanHeightNcdRisc2017Dataset, error) {
	var response model.HumanHeightNcdRisc2017Dataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) HumanHeightUniversityOfTuebingen2015(ctx context.Context, entity string, year int) (*model.HumanHeightUniversityOfTuebingen2015Dataset, error) {
	var response model.HumanHeightUniversityOfTuebingen2015Dataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) HumanRightsProtectionScoreChristopherFarris2014AndKeithSchnakenberg(ctx context.Context, entity string, year int) (*model.HumanRightsProtectionScoreChristopherFarris2014AndKeithSchnakenbergDataset, error) {
	var response model.HumanRightsProtectionScoreChristopherFarris2014AndKeithSchnakenbergDataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) HumanRightsProtectionFarissEtAl2020(ctx context.Context, entity string, year int) (*model.HumanRightsProtectionFarissEtAl2020Dataset, error) {
	var response model.HumanRightsProtectionFarissEtAl2020Dataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) HurricaneForecastingErrorNhc2019(ctx context.Context, entity string, year int) (*model.HurricaneForecastingErrorNhc2019Dataset, error) {
	var response model.HurricaneForecastingErrorNhc2019Dataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) HurricaneLandfallsContinentalUsHurdatNoaa(ctx context.Context, entity string, year int) (*model.HurricaneLandfallsContinentalUsHurdatNoaaDataset, error) {
	var response model.HurricaneLandfallsContinentalUsHurdatNoaaDataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) HypotheticalGlobalCo2EmissionsCdiac2014(ctx context.Context, entity string, year int) (*model.HypotheticalGlobalCo2EmissionsCdiac2014Dataset, error) {
	var response model.HypotheticalGlobalCo2EmissionsCdiac2014Dataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) HypotheticalMeatConsumptionOwidBasedOnFaoAndUn(ctx context.Context, entity string, year int) (*model.HypotheticalMeatConsumptionOwidBasedOnFaoAndUnDataset, error) {
	var response model.HypotheticalMeatConsumptionOwidBasedOnFaoAndUnDataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) IpccScenariosIiasa(ctx context.Context, entity string, year int) (*model.IpccScenariosIiasaDataset, error) {
	var response model.IpccScenariosIiasaDataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) IqDataPietschnigAndVoracek2015(ctx context.Context, entity string, year int) (*model.IqDataPietschnigAndVoracek2015Dataset, error) {
	var response model.IqDataPietschnigAndVoracek2015Dataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) IncidenceOfChildLaborEnglandItalyUsWorldCunninghamAndViazzo1996AndOthers(ctx context.Context, entity string, year int) (*model.IncidenceOfChildLaborEnglandItalyUsWorldCunninghamAndViazzo1996AndOthersDataset, error) {
	var response model.IncidenceOfChildLaborEnglandItalyUsWorldCunninghamAndViazzo1996AndOthersDataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) IncidenceOfManagerialOrProfessionalJobsAndCollectiveBargainingByGenderBlauAndKahn2017(ctx context.Context, entity string, year int) (*model.IncidenceOfManagerialOrProfessionalJobsAndCollectiveBargainingByGenderBlauAndKahn2017Dataset, error) {
	var response model.IncidenceOfManagerialOrProfessionalJobsAndCollectiveBargainingByGenderBlauAndKahn2017Dataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) IncomeClassificationWorldBank2017(ctx context.Context, entity string, year int) (*model.IncomeClassificationWorldBank2017Dataset, error) {
	var response model.IncomeClassificationWorldBank2017Dataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) IncomesAcrossTheDistributionDatabaseNolanThewissenRoserInLevels2016(ctx context.Context, entity string, year int) (*model.IncomesAcrossTheDistributionDatabaseNolanThewissenRoserInLevels2016Dataset, error) {
	var response model.IncomesAcrossTheDistributionDatabaseNolanThewissenRoserInLevels2016Dataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) IncomesAcrossTheDistributionDatabaseAuthoredByNolanThewissenRoserBasedOnLisIndexedToTheFirstYear2016(ctx context.Context, entity string, year int) (*model.IncomesAcrossTheDistributionDatabaseAuthoredByNolanThewissenRoserBasedOnLisIndexedToTheFirstYear2016Dataset, error) {
	var response model.IncomesAcrossTheDistributionDatabaseAuthoredByNolanThewissenRoserBasedOnLisIndexedToTheFirstYear2016Dataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) IncomesAcrossTheDistributionDatabaseGini2016(ctx context.Context, entity string, year int) (*model.IncomesAcrossTheDistributionDatabaseGini2016Dataset, error) {
	var response model.IncomesAcrossTheDistributionDatabaseGini2016Dataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) IndicatorsForWhatIsPppWorldBank(ctx context.Context, entity string, year int) (*model.IndicatorsForWhatIsPppWorldBankDataset, error) {
	var response model.IndicatorsForWhatIsPppWorldBankDataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) IndustrialMotivePowerInTheUk1800_70Musson1976(ctx context.Context, entity string, year int) (*model.IndustrialMotivePowerInTheUk180070Musson1976Dataset, error) {
	var response model.IndustrialMotivePowerInTheUk180070Musson1976Dataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) InequalityBeforeAndAfterTaxesOecd2008(ctx context.Context, entity string, year int) (*model.InequalityBeforeAndAfterTaxesOecd2008Dataset, error) {
	var response model.InequalityBeforeAndAfterTaxesOecd2008Dataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) InequalityInLatinAmericaSedlacCedlasAndTheWorldBank(ctx context.Context, entity string, year int) (*model.InequalityInLatinAmericaSedlacCedlasAndTheWorldBankDataset, error) {
	var response model.InequalityInLatinAmericaSedlacCedlasAndTheWorldBankDataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) InequalityInHumanDevelopmentIndicesUndp2019(ctx context.Context, entity string, year int) (*model.InequalityInHumanDevelopmentIndicesUndp2019Dataset, error) {
	var response model.InequalityInHumanDevelopmentIndicesUndp2019Dataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) InfantMortalityRateIhme2017(ctx context.Context, entity string, year int) (*model.InfantMortalityRateIhme2017Dataset, error) {
	var response model.InfantMortalityRateIhme2017Dataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) InheritanceForWomenHowWasLife2014(ctx context.Context, entity string, year int) (*model.InheritanceForWomenHowWasLife2014Dataset, error) {
	var response model.InheritanceForWomenHowWasLife2014Dataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) IntegratedNetworkForSocietalConflictResearchPoliticalInstabilityTaskForcePitf(ctx context.Context, entity string, year int) (*model.IntegratedNetworkForSocietalConflictResearchPoliticalInstabilityTaskForcePitfDataset, error) {
	var response model.IntegratedNetworkForSocietalConflictResearchPoliticalInstabilityTaskForcePitfDataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) IntercontinentalTradeCostaPalmaAndReis2015(ctx context.Context, entity string, year int) (*model.IntercontinentalTradeCostaPalmaAndReis2015Dataset, error) {
	var response model.IntercontinentalTradeCostaPalmaAndReis2015Dataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) InternationalHistoricalStatisticsBirthsPer1_000BrianMitchell2013(ctx context.Context, entity string, year int) (*model.InternationalHistoricalStatisticsBirthsPer1000BrianMitchell2013Dataset, error) {
	var response model.InternationalHistoricalStatisticsBirthsPer1000BrianMitchell2013Dataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) InternationalHistoricalStatisticsDeathsPer1_000BrianMitchell2013(ctx context.Context, entity string, year int) (*model.InternationalHistoricalStatisticsDeathsPer1000BrianMitchell2013Dataset, error) {
	var response model.InternationalHistoricalStatisticsDeathsPer1000BrianMitchell2013Dataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) InternationalHistoricalStatisticsEuropeanTradeBrianMitchell2015(ctx context.Context, entity string, year int) (*model.InternationalHistoricalStatisticsEuropeanTradeBrianMitchell2015Dataset, error) {
	var response model.InternationalHistoricalStatisticsEuropeanTradeBrianMitchell2015Dataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) InterpersonalTrustGeneralSocialSurveyGss(ctx context.Context, entity string, year int) (*model.InterpersonalTrustGeneralSocialSurveyGssDataset, error) {
	var response model.InterpersonalTrustGeneralSocialSurveyGssDataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) InvestmentInRenewablesByRegionIrena2016(ctx context.Context, entity string, year int) (*model.InvestmentInRenewablesByRegionIrena2016Dataset, error) {
	var response model.InvestmentInRenewablesByRegionIrena2016Dataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) InvestmentInRenewablesByTechnologyIrena2017(ctx context.Context, entity string, year int) (*model.InvestmentInRenewablesByTechnologyIrena2017Dataset, error) {
	var response model.InvestmentInRenewablesByTechnologyIrena2017Dataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) JobSearchMethodsUsPewResearchCenter2015(ctx context.Context, entity string, year int) (*model.JobSearchMethodsUsPewResearchCenter2015Dataset, error) {
	var response model.JobSearchMethodsUsPewResearchCenter2015Dataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) LgbtMaritalStatusInTheUsGallup2017(ctx context.Context, entity string, year int) (*model.LgbtMaritalStatusInTheUsGallup2017Dataset, error) {
	var response model.LgbtMaritalStatusInTheUsGallup2017Dataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) LisKeyFiguresLuxumbourgIncomeStudy(ctx context.Context, entity string, year int) (*model.LisKeyFiguresLuxumbourgIncomeStudyDataset, error) {
	var response model.LisKeyFiguresLuxumbourgIncomeStudyDataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) LabGrownMeatPricesNextbigfuture2017AndUnitedStatesBureauOfLaborStatisticsBls2017(ctx context.Context, entity string, year int) (*model.LabGrownMeatPricesNextbigfuture2017AndUnitedStatesBureauOfLaborStatisticsBls2017Dataset, error) {
	var response model.LabGrownMeatPricesNextbigfuture2017AndUnitedStatesBureauOfLaborStatisticsBls2017Dataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) LaborForceParticipationRatesOfMenAge65AndOverInTheUsOwidBasedOnShort2002AndOecd(ctx context.Context, entity string, year int) (*model.LaborForceParticipationRatesOfMenAge65AndOverInTheUsOwidBasedOnShort2002AndOecdDataset, error) {
	var response model.LaborForceParticipationRatesOfMenAge65AndOverInTheUsOwidBasedOnShort2002AndOecdDataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) LaborProductivityInCottonSpinningChapman1972(ctx context.Context, entity string, year int) (*model.LaborProductivityInCottonSpinningChapman1972Dataset, error) {
	var response model.LaborProductivityInCottonSpinningChapman1972Dataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) LaborProductivityInCottonSpinningAndWeavingEllison1886(ctx context.Context, entity string, year int) (*model.LaborProductivityInCottonSpinningAndWeavingEllison1886Dataset, error) {
	var response model.LaborProductivityInCottonSpinningAndWeavingEllison1886Dataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) LaborProductivityPerHourHillThomasDimsdale2016BankOfEngland(ctx context.Context, entity string, year int) (*model.LaborProductivityPerHourHillThomasDimsdale2016BankOfEnglandDataset, error) {
	var response model.LaborProductivityPerHourHillThomasDimsdale2016BankOfEnglandDataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) LabourCostRatio45_54YearOldPopulation2009Oecd2012(ctx context.Context, entity string, year int) (*model.LabourCostRatio4554YearOldPopulation2009Oecd2012Dataset, error) {
	var response model.LabourCostRatio4554YearOldPopulation2009Oecd2012Dataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) LandUseDataHyde2017(ctx context.Context, entity string, year int) (*model.LandUseDataHyde2017Dataset, error) {
	var response model.LandUseDataHyde2017Dataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) LandUseMapByAreaOwidBasedOnFao(ctx context.Context, entity string, year int) (*model.LandUseMapByAreaOwidBasedOnFaoDataset, error) {
	var response model.LandUseMapByAreaOwidBasedOnFaoDataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) LandUnderCerealProductionIndexWorldBank2017(ctx context.Context, entity string, year int) (*model.LandUnderCerealProductionIndexWorldBank2017Dataset, error) {
	var response model.LandUnderCerealProductionIndexWorldBank2017Dataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) LandUseSince10000bcEllisEtAl2020(ctx context.Context, entity string, year int) (*model.LandUseSince10000bcEllisEtAl2020Dataset, error) {
	var response model.LandUseSince10000bcEllisEtAl2020Dataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) LargestCitiesByPopulationDensityUnHabitat2014(ctx context.Context, entity string, year int) (*model.LargestCitiesByPopulationDensityUnHabitat2014Dataset, error) {
	var response model.LargestCitiesByPopulationDensityUnHabitat2014Dataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) LearningCostsJDoyneFarmerAndFrancoisLafond2016(ctx context.Context, entity string, year int) (*model.LearningCostsJDoyneFarmerAndFrancoisLafond2016Dataset, error) {
	var response model.LearningCostsJDoyneFarmerAndFrancoisLafond2016Dataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) LearningAdjustedYearsOfSchoolingWorldBank2018(ctx context.Context, entity string, year int) (*model.LearningAdjustedYearsOfSchoolingWorldBank2018Dataset, error) {
	var response model.LearningAdjustedYearsOfSchoolingWorldBank2018Dataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) LengthOfTheWorkDayFrom1890sTo1991Costa2000(ctx context.Context, entity string, year int) (*model.LengthOfTheWorkDayFrom1890sTo1991Costa2000Dataset, error) {
	var response model.LengthOfTheWorkDayFrom1890sTo1991Costa2000Dataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) LengthOfTheWorkdayIn1880AtackAndBateman1992(ctx context.Context, entity string, year int) (*model.LengthOfTheWorkdayIn1880AtackAndBateman1992Dataset, error) {
	var response model.LengthOfTheWorkdayIn1880AtackAndBateman1992Dataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) LevelsOfUrbanizationAndPerCapitaGnpInVariousRegionsBairoch1988(ctx context.Context, entity string, year int) (*model.LevelsOfUrbanizationAndPerCapitaGnpInVariousRegionsBairoch1988Dataset, error) {
	var response model.LevelsOfUrbanizationAndPerCapitaGnpInVariousRegionsBairoch1988Dataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) LifeExpectancy1950_2015UnPopulationDivision2015(ctx context.Context, entity string, year int) (*model.LifeExpectancy19502015UnPopulationDivision2015Dataset, error) {
	var response model.LifeExpectancy19502015UnPopulationDivision2015Dataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) LifeExpectancyAtBirthWorldBank2015(ctx context.Context, entity string, year int) (*model.LifeExpectancyAtBirthWorldBank2015Dataset, error) {
	var response model.LifeExpectancyAtBirthWorldBank2015Dataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) LifeExpectancyAtBirthBothGendersClioInfra(ctx context.Context, entity string, year int) (*model.LifeExpectancyAtBirthBothGendersClioInfraDataset, error) {
	var response model.LifeExpectancyAtBirthBothGendersClioInfraDataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) LifeExpectancyGapminderUnIhme(ctx context.Context, entity string, year int) (*model.LifeExpectancyGapminderUnIhmeDataset, error) {
	var response model.LifeExpectancyGapminderUnIhmeDataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) LifeExpectancyOecd(ctx context.Context, entity string, year int) (*model.LifeExpectancyOecdDataset, error) {
	var response model.LifeExpectancyOecdDataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) LifeExpectancyRiley2005AndUn(ctx context.Context, entity string, year int) (*model.LifeExpectancyRiley2005AndUnDataset, error) {
	var response model.LifeExpectancyRiley2005AndUnDataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) LifeExpectancyRiley2005ClioInfra2015AndUn2019(ctx context.Context, entity string, year int) (*model.LifeExpectancyRiley2005ClioInfra2015AndUn2019Dataset, error) {
	var response model.LifeExpectancyRiley2005ClioInfra2015AndUn2019Dataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) LifeExpectancyAtAge10_1750_2100UnitedNationsPopulationDivisionAndHumanMortalityDatabase2015(ctx context.Context, entity string, year int) (*model.LifeExpectancyAtAge1017502100UnitedNationsPopulationDivisionAndHumanMortalityDatabase2015Dataset, error) {
	var response model.LifeExpectancyAtAge1017502100UnitedNationsPopulationDivisionAndHumanMortalityDatabase2015Dataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) LifeExpectancyProjectionsUkOns(ctx context.Context, entity string, year int) (*model.LifeExpectancyProjectionsUkOnsDataset, error) {
	var response model.LifeExpectancyProjectionsUkOnsDataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) LifeExpectancyJamesRileyForData1990AndEarlierWhoAndWorldBankForLaterDataByMaxRoser(ctx context.Context, entity string, year int) (*model.LifeExpectancyJamesRileyForData1990AndEarlierWhoAndWorldBankForLaterDataByMaxRoserDataset, error) {
	var response model.LifeExpectancyJamesRileyForData1990AndEarlierWhoAndWorldBankForLaterDataByMaxRoserDataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) LifeExpectationBySexAtAges0_15And45OwidBasedOnHacker2010AndTheUsSocialSecurityAdministration2017(ctx context.Context, entity string, year int) (*model.LifeExpectationBySexAtAges015And45OwidBasedOnHacker2010AndTheUsSocialSecurityAdministration2017Dataset, error) {
	var response model.LifeExpectationBySexAtAges015And45OwidBasedOnHacker2010AndTheUsSocialSecurityAdministration2017Dataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) LifeSatisfactionEurobarometer2017(ctx context.Context, entity string, year int) (*model.LifeSatisfactionEurobarometer2017Dataset, error) {
	var response model.LifeSatisfactionEurobarometer2017Dataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) LifeSatisfactionWorldValueSurvey2014(ctx context.Context, entity string, year int) (*model.LifeSatisfactionWorldValueSurvey2014Dataset, error) {
	var response model.LifeSatisfactionWorldValueSurvey2014Dataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) LifeCycleImpactsOfEnergySourcesUnece(ctx context.Context, entity string, year int) (*model.LifeCycleImpactsOfEnergySourcesUneceDataset, error) {
	var response model.LifeCycleImpactsOfEnergySourcesUneceDataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) LightingEfficiencyInUkOwidBasedOnFouquetAndPearson2007(ctx context.Context, entity string, year int) (*model.LightingEfficiencyInUkOwidBasedOnFouquetAndPearson2007Dataset, error) {
	var response model.LightingEfficiencyInUkOwidBasedOnFouquetAndPearson2007Dataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) LiteracyByYearsOfSchoolingUs1947Oecd2014(ctx context.Context, entity string, year int) (*model.LiteracyByYearsOfSchoolingUs1947Oecd2014Dataset, error) {
	var response model.LiteracyByYearsOfSchoolingUs1947Oecd2014Dataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) LiteracyInEnglandBySexSchofield1973Houston1982Cressy1980(ctx context.Context, entity string, year int) (*model.LiteracyInEnglandBySexSchofield1973Houston1982Cressy1980Dataset, error) {
	var response model.LiteracyInEnglandBySexSchofield1973Houston1982Cressy1980Dataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) LiteracyRatePercOfTotalRespondentsDhsSurveys(ctx context.Context, entity string, year int) (*model.LiteracyRatePercOfTotalRespondentsDhsSurveysDataset, error) {
	var response model.LiteracyRatePercOfTotalRespondentsDhsSurveysDataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) LiterateWorldPopulationOurworldindataBasedOnOecdAndUnesco(ctx context.Context, entity string, year int) (*model.LiterateWorldPopulationOurworldindataBasedOnOecdAndUnescoDataset, error) {
	var response model.LiterateWorldPopulationOurworldindataBasedOnOecdAndUnescoDataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) LivestockCountsHydeAndFao2017(ctx context.Context, entity string, year int) (*model.LivestockCountsHydeAndFao2017Dataset, error) {
	var response model.LivestockCountsHydeAndFao2017Dataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) LivingPlanetIndexWwf2020(ctx context.Context, entity string, year int) (*model.LivingPlanetIndexWwf2020Dataset, error) {
	var response model.LivingPlanetIndexWwf2020Dataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) LongRunLifeExpectancyGapminderUn(ctx context.Context, entity string, year int) (*model.LongRunLifeExpectancyGapminderUnDataset, error) {
	var response model.LongRunLifeExpectancyGapminderUnDataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) LongRunTimeUseInNorwayBySexNorwayStatistics(ctx context.Context, entity string, year int) (*model.LongRunTimeUseInNorwayBySexNorwayStatisticsDataset, error) {
	var response model.LongRunTimeUseInNorwayBySexNorwayStatisticsDataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) LongTermProductivityBergeaudCetteAndLecat2016(ctx context.Context, entity string, year int) (*model.LongTermProductivityBergeaudCetteAndLecat2016Dataset, error) {
	var response model.LongTermProductivityBergeaudCetteAndLecat2016Dataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) LongRunSeriesOfHealthExpenditureWorldBankWdi2017(ctx context.Context, entity string, year int) (*model.LongRunSeriesOfHealthExpenditureWorldBankWdi2017Dataset, error) {
	var response model.LongRunSeriesOfHealthExpenditureWorldBankWdi2017Dataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) LongTermEnergyTransitionInEuropeGalesEtAl2007(ctx context.Context, entity string, year int) (*model.LongTermEnergyTransitionInEuropeGalesEtAl2007Dataset, error) {
	var response model.LongTermEnergyTransitionInEuropeGalesEtAl2007Dataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) LongTermEnergyTransitionsEnergyHistoryHarvard2016(ctx context.Context, entity string, year int) (*model.LongTermEnergyTransitionsEnergyHistoryHarvard2016Dataset, error) {
	var response model.LongTermEnergyTransitionsEnergyHistoryHarvard2016Dataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) LongTermPerCapitaFossilFuelsOwidBasedOnUnGapminderBpEtemadAndLuciana(ctx context.Context, entity string, year int) (*model.LongTermPerCapitaFossilFuelsOwidBasedOnUnGapminderBpEtemadAndLucianaDataset, error) {
	var response model.LongTermPerCapitaFossilFuelsOwidBasedOnUnGapminderBpEtemadAndLucianaDataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) LongTermWheatYieldsFao2017AndBaylissSmith1984(ctx context.Context, entity string, year int) (*model.LongTermWheatYieldsFao2017AndBaylissSmith1984Dataset, error) {
	var response model.LongTermWheatYieldsFao2017AndBaylissSmith1984Dataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) LongTermYieldsInTheUnitedKingdom2022(ctx context.Context, entity string, year int) (*model.LongTermYieldsInTheUnitedKingdom2022Dataset, error) {
	var response model.LongTermYieldsInTheUnitedKingdom2022Dataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) LostSchoolGrantsReinikkaAndSvensson2004(ctx context.Context, entity string, year int) (*model.LostSchoolGrantsReinikkaAndSvensson2004Dataset, error) {
	var response model.LostSchoolGrantsReinikkaAndSvensson2004Dataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) LowestPayingOccupationsPercentFemaleNwlc2014(ctx context.Context, entity string, year int) (*model.LowestPayingOccupationsPercentFemaleNwlc2014Dataset, error) {
	var response model.LowestPayingOccupationsPercentFemaleNwlc2014Dataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) LungCancerMortalityRatesPer100_000_2022(ctx context.Context, entity string, year int) (*model.LungCancerMortalityRatesPer1000002022Dataset, error) {
	var response model.LungCancerMortalityRatesPer1000002022Dataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) MdgFinalEvaluationUnMdgReport(ctx context.Context, entity string, year int) (*model.MdgFinalEvaluationUnMdgReportDataset, error) {
	var response model.MdgFinalEvaluationUnMdgReportDataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) MaddisonProjectDatabase2018BoltEtAl2018(ctx context.Context, entity string, year int) (*model.MaddisonProjectDatabase2018BoltEtAl2018Dataset, error) {
	var response model.MaddisonProjectDatabase2018BoltEtAl2018Dataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) MaddisonProjectDatabase2020BoltAndVanZanden2020(ctx context.Context, entity string, year int) (*model.MaddisonProjectDatabase2020BoltAndVanZanden2020Dataset, error) {
	var response model.MaddisonProjectDatabase2020BoltAndVanZanden2020Dataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) MalariaDeathsIhme2016(ctx context.Context, entity string, year int) (*model.MalariaDeathsIhme2016Dataset, error) {
	var response model.MalariaDeathsIhme2016Dataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) MaleAndFemaleLifeExpectancyByAgeInTheLongRunHumanMortalityDatabase2018AndOthers(ctx context.Context, entity string, year int) (*model.MaleAndFemaleLifeExpectancyByAgeInTheLongRunHumanMortalityDatabase2018AndOthersDataset, error) {
	var response model.MaleAndFemaleLifeExpectancyByAgeInTheLongRunHumanMortalityDatabase2018AndOthersDataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) MaleToFemaleRatioHighSchoolCoursesInUsaGoldinEtAl(ctx context.Context, entity string, year int) (*model.MaleToFemaleRatioHighSchoolCoursesInUsaGoldinEtAlDataset, error) {
	var response model.MaleToFemaleRatioHighSchoolCoursesInUsaGoldinEtAlDataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) MarineEnergyIrena(ctx context.Context, entity string, year int) (*model.MarineEnergyIrenaDataset, error) {
	var response model.MarineEnergyIrenaDataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) MarineStocksByRegionAndTaxaRamlegacy(ctx context.Context, entity string, year int) (*model.MarineStocksByRegionAndTaxaRamlegacyDataset, error) {
	var response model.MarineStocksByRegionAndTaxaRamlegacyDataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) MarketShareOfIodizedSaltInEuropeanCountriesEuropeanCommission2006(ctx context.Context, entity string, year int) (*model.MarketShareOfIodizedSaltInEuropeanCountriesEuropeanCommission2006Dataset, error) {
	var response model.MarketShareOfIodizedSaltInEuropeanCountriesEuropeanCommission2006Dataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) MaternalMortalityRatioGapminder2010AndWorldBank2015(ctx context.Context, entity string, year int) (*model.MaternalMortalityRatioGapminder2010AndWorldBank2015Dataset, error) {
	var response model.MaternalMortalityRatioGapminder2010AndWorldBank2015Dataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) MaternalDeathsTo2030BauVsSdgTargetBasedOnWorldBankAndUn2018(ctx context.Context, entity string, year int) (*model.MaternalDeathsTo2030BauVsSdgTargetBasedOnWorldBankAndUn2018Dataset, error) {
	var response model.MaternalDeathsTo2030BauVsSdgTargetBasedOnWorldBankAndUn2018Dataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) MaternalMortalityProjectionTo2030BasedOnWorldBank2018(ctx context.Context, entity string, year int) (*model.MaternalMortalityProjectionTo2030BasedOnWorldBank2018Dataset, error) {
	var response model.MaternalMortalityProjectionTo2030BasedOnWorldBank2018Dataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) MeanBmiNcdRisc2017(ctx context.Context, entity string, year int) (*model.MeanBmiNcdRisc2017Dataset, error) {
	var response model.MeanBmiNcdRisc2017Dataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) MeanYearsOfSchoolingWomen15To49OurWorldInData2017(ctx context.Context, entity string, year int) (*model.MeanYearsOfSchoolingWomen15To49OurWorldInData2017Dataset, error) {
	var response model.MeanYearsOfSchoolingWomen15To49OurWorldInData2017Dataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) MeaslesLondon(ctx context.Context, entity string, year int) (*model.MeaslesLondonDataset, error) {
	var response model.MeaslesLondonDataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) MeasuresAndIndicatorsForPovertyPovcalnetWorldBank2017(ctx context.Context, entity string, year int) (*model.MeasuresAndIndicatorsForPovertyPovcalnetWorldBank2017Dataset, error) {
	var response model.MeasuresAndIndicatorsForPovertyPovcalnetWorldBank2017Dataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) MeatConsumptionInEu28Oecd2018(ctx context.Context, entity string, year int) (*model.MeatConsumptionInEu28Oecd2018Dataset, error) {
	var response model.MeatConsumptionInEu28Oecd2018Dataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) MeatConsumptionInTheUsaUsda2018(ctx context.Context, entity string, year int) (*model.MeatConsumptionInTheUsaUsda2018Dataset, error) {
	var response model.MeatConsumptionInTheUsaUsda2018Dataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) MeatConversionEfficienciesAlexanderEtAl2016(ctx context.Context, entity string, year int) (*model.MeatConversionEfficienciesAlexanderEtAl2016Dataset, error) {
	var response model.MeatConversionEfficienciesAlexanderEtAl2016Dataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) MedianUnPopulationProjectionsGlobalVsAfricaOwidBasedOnUn(ctx context.Context, entity string, year int) (*model.MedianUnPopulationProjectionsGlobalVsAfricaOwidBasedOnUnDataset, error) {
	var response model.MedianUnPopulationProjectionsGlobalVsAfricaOwidBasedOnUnDataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) MentalAndSubstanceUseDisorderDisaggregatedIhme(ctx context.Context, entity string, year int) (*model.MentalAndSubstanceUseDisorderDisaggregatedIhmeDataset, error) {
	var response model.MentalAndSubstanceUseDisorderDisaggregatedIhmeDataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) MentalHealthAsRiskFactorForSubstanceUseSwendsenEtAl2010(ctx context.Context, entity string, year int) (*model.MentalHealthAsRiskFactorForSubstanceUseSwendsenEtAl2010Dataset, error) {
	var response model.MentalHealthAsRiskFactorForSubstanceUseSwendsenEtAl2010Dataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) MentalHealthServicesAcrossIncomesWangEtAl2007(ctx context.Context, entity string, year int) (*model.MentalHealthServicesAcrossIncomesWangEtAl2007Dataset, error) {
	var response model.MentalHealthServicesAcrossIncomesWangEtAl2007Dataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) MetalProductionClioInfraAndUsgs(ctx context.Context, entity string, year int) (*model.MetalProductionClioInfraAndUsgsDataset, error) {
	var response model.MetalProductionClioInfraAndUsgsDataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) MethaneEmissionsBySectorCait2020(ctx context.Context, entity string, year int) (*model.MethaneEmissionsBySectorCait2020Dataset, error) {
	var response model.MethaneEmissionsBySectorCait2020Dataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) MethaneEmissionsBySectorCait2021(ctx context.Context, entity string, year int) (*model.MethaneEmissionsBySectorCait2021Dataset, error) {
	var response model.MethaneEmissionsBySectorCait2021Dataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) MilestonesOfWomensPoliticalRepresentationPaxtonEtAl2006(ctx context.Context, entity string, year int) (*model.MilestonesOfWomensPoliticalRepresentationPaxtonEtAl2006Dataset, error) {
	var response model.MilestonesOfWomensPoliticalRepresentationPaxtonEtAl2006Dataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) MilitaryExpenditureAsAShareOfGdpOwidBasedOnCowAndSipri2017(ctx context.Context, entity string, year int) (*model.MilitaryExpenditureAsAShareOfGdpOwidBasedOnCowAndSipri2017Dataset, error) {
	var response model.MilitaryExpenditureAsAShareOfGdpOwidBasedOnCowAndSipri2017Dataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) MineralProductionBgs2016(ctx context.Context, entity string, year int) (*model.MineralProductionBgs2016Dataset, error) {
	var response model.MineralProductionBgs2016Dataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) MinimumReadingAndMathsProficiencyGemReport20178(ctx context.Context, entity string, year int) (*model.MinimumReadingAndMathsProficiencyGemReport20178Dataset, error) {
	var response model.MinimumReadingAndMathsProficiencyGemReport20178Dataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) MissingPlasticBudgetLebretonEtAl2019(ctx context.Context, entity string, year int) (*model.MissingPlasticBudgetLebretonEtAl2019Dataset, error) {
	var response model.MissingPlasticBudgetLebretonEtAl2019Dataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) MissingWomenEstimatesBongaartsAndGuilmoto2015(ctx context.Context, entity string, year int) (*model.MissingWomenEstimatesBongaartsAndGuilmoto2015Dataset, error) {
	var response model.MissingWomenEstimatesBongaartsAndGuilmoto2015Dataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) MobileBankAccountsByRegionGsma2019(ctx context.Context, entity string, year int) (*model.MobileBankAccountsByRegionGsma2019Dataset, error) {
	var response model.MobileBankAccountsByRegionGsma2019Dataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) MortalityFromAllFormsOfViolenceIhme2016(ctx context.Context, entity string, year int) (*model.MortalityFromAllFormsOfViolenceIhme2016Dataset, error) {
	var response model.MortalityFromAllFormsOfViolenceIhme2016Dataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) MotorVehiclesPer1000PeopleNationmaster2014(ctx context.Context, entity string, year int) (*model.MotorVehiclesPer1000PeopleNationmaster2014Dataset, error) {
	var response model.MotorVehiclesPer1000PeopleNationmaster2014Dataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) MultinationalTimeUseStudyMtusGershunyAndFisher2013(ctx context.Context, entity string, year int) (*model.MultinationalTimeUseStudyMtusGershunyAndFisher2013Dataset, error) {
	var response model.MultinationalTimeUseStudyMtusGershunyAndFisher2013Dataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) NihDnaSequencingCosts(ctx context.Context, entity string, year int) (*model.NihDnaSequencingCostsDataset, error) {
	var response model.NihDnaSequencingCostsDataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) NationalPovertyLinesJolliffeAndPrydz2016(ctx context.Context, entity string, year int) (*model.NationalPovertyLinesJolliffeAndPrydz2016Dataset, error) {
	var response model.NationalPovertyLinesJolliffeAndPrydz2016Dataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) NaturalDisastersEmdatDecadal(ctx context.Context, entity string, year int) (*model.NaturalDisastersEmdatDecadalDataset, error) {
	var response model.NaturalDisastersEmdatDecadalDataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) NaturalDisastersEmdat(ctx context.Context, entity string, year int) (*model.NaturalDisastersEmdatDataset, error) {
	var response model.NaturalDisastersEmdatDataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) NaturalDisastersFrom1900To2019Emdat2020(ctx context.Context, entity string, year int) (*model.NaturalDisastersFrom1900To2019Emdat2020Dataset, error) {
	var response model.NaturalDisastersFrom1900To2019Emdat2020Dataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) NeglectedTropicalDiseasesLymphaticFilariasisPopulationRequiringPcNotTreatedAndTreatedEnricJane2016(ctx context.Context, entity string, year int) (*model.NeglectedTropicalDiseasesLymphaticFilariasisPopulationRequiringPcNotTreatedAndTreatedEnricJane2016Dataset, error) {
	var response model.NeglectedTropicalDiseasesLymphaticFilariasisPopulationRequiringPcNotTreatedAndTreatedEnricJane2016Dataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) NeonatalMortalityRateViaChildmortalityorg2015(ctx context.Context, entity string, year int) (*model.NeonatalMortalityRateViaChildmortalityorg2015Dataset, error) {
	var response model.NeonatalMortalityRateViaChildmortalityorg2015Dataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) NeonatalTetanusIncidence(ctx context.Context, entity string, year int) (*model.NeonatalTetanusIncidenceDataset, error) {
	var response model.NeonatalTetanusIncidenceDataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) NewEstimatesOfHoursOfWorkPerWeek1900_1957Jones1963(ctx context.Context, entity string, year int) (*model.NewEstimatesOfHoursOfWorkPerWeek19001957Jones1963Dataset, error) {
	var response model.NewEstimatesOfHoursOfWorkPerWeek19001957Jones1963Dataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) NewsworthinessOfDisastersByDisasterTypeAndRegionEisenseeAndStromberg2007(ctx context.Context, entity string, year int) (*model.NewsworthinessOfDisastersByDisasterTypeAndRegionEisenseeAndStromberg2007Dataset, error) {
	var response model.NewsworthinessOfDisastersByDisasterTypeAndRegionEisenseeAndStromberg2007Dataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) NitrogenFertilizerConsumptionFao2017(ctx context.Context, entity string, year int) (*model.NitrogenFertilizerConsumptionFao2017Dataset, error) {
	var response model.NitrogenFertilizerConsumptionFao2017Dataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) NitrogenFertilizerProductionFao2017(ctx context.Context, entity string, year int) (*model.NitrogenFertilizerProductionFao2017Dataset, error) {
	var response model.NitrogenFertilizerProductionFao2017Dataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) NitrousOxideEmissionsBySectorCait2020(ctx context.Context, entity string, year int) (*model.NitrousOxideEmissionsBySectorCait2020Dataset, error) {
	var response model.NitrousOxideEmissionsBySectorCait2020Dataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) NitrousOxideEmissionsBySectorCait2021(ctx context.Context, entity string, year int) (*model.NitrousOxideEmissionsBySectorCait2021Dataset, error) {
	var response model.NitrousOxideEmissionsBySectorCait2021Dataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) NonCommercialFlightDistanceRecordsWikipedia(ctx context.Context, entity string, year int) (*model.NonCommercialFlightDistanceRecordsWikipediaDataset, error) {
	var response model.NonCommercialFlightDistanceRecordsWikipediaDataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) NorthAtlanticHurricanesHudratNoaa(ctx context.Context, entity string, year int) (*model.NorthAtlanticHurricanesHudratNoaaDataset, error) {
	var response model.NorthAtlanticHurricanesHudratNoaaDataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) NuclearWeaponsTestsArmsControlAssociation2020(ctx context.Context, entity string, year int) (*model.NuclearWeaponsTestsArmsControlAssociation2020Dataset, error) {
	var response model.NuclearWeaponsTestsArmsControlAssociation2020Dataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) NuclearWarheadStockpilesFederationOfAmericanScientists(ctx context.Context, entity string, year int) (*model.NuclearWarheadStockpilesFederationOfAmericanScientistsDataset, error) {
	var response model.NuclearWarheadStockpilesFederationOfAmericanScientistsDataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) NuclearWeaponsProliferationTotalOwidBasedOnBleek2017NuclearThreatInitiative2022(ctx context.Context, entity string, year int) (*model.NuclearWeaponsProliferationTotalOwidBasedOnBleek2017NuclearThreatInitiative2022Dataset, error) {
	var response model.NuclearWeaponsProliferationTotalOwidBasedOnBleek2017NuclearThreatInitiative2022Dataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) NuclearWeaponsProliferationOwidBasedOnBleek2017NuclearThreatInitiative2022(ctx context.Context, entity string, year int) (*model.NuclearWeaponsProliferationOwidBasedOnBleek2017NuclearThreatInitiative2022Dataset, error) {
	var response model.NuclearWeaponsProliferationOwidBasedOnBleek2017NuclearThreatInitiative2022Dataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) NumberAndPercentageOfCurrentSmokersBySexAmericanLungAssociation2011(ctx context.Context, entity string, year int) (*model.NumberAndPercentageOfCurrentSmokersBySexAmericanLungAssociation2011Dataset, error) {
	var response model.NumberAndPercentageOfCurrentSmokersBySexAmericanLungAssociation2011Dataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) NumberOfChildDeaths1950_2017Ihme2017(ctx context.Context, entity string, year int) (*model.NumberOfChildDeaths19502017Ihme2017Dataset, error) {
	var response model.NumberOfChildDeaths19502017Ihme2017Dataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) NumberOfChildrenWhoAreStuntedOwidBasedOnUnicefwho(ctx context.Context, entity string, year int) (*model.NumberOfChildrenWhoAreStuntedOwidBasedOnUnicefwhoDataset, error) {
	var response model.NumberOfChildrenWhoAreStuntedOwidBasedOnUnicefwhoDataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) NumberOfCountriesWithMinimumUrbanPopulationThresholdUn2018(ctx context.Context, entity string, year int) (*model.NumberOfCountriesWithMinimumUrbanPopulationThresholdUn2018Dataset, error) {
	var response model.NumberOfCountriesWithMinimumUrbanPopulationThresholdUn2018Dataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) NumberOfDeathsDueToTetanus(ctx context.Context, entity string, year int) (*model.NumberOfDeathsDueToTetanusDataset, error) {
	var response model.NumberOfDeathsDueToTetanusDataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) NumberOfDeathsInEnglandAndWalesByAgeOns(ctx context.Context, entity string, year int) (*model.NumberOfDeathsInEnglandAndWalesByAgeOnsDataset, error) {
	var response model.NumberOfDeathsInEnglandAndWalesByAgeOnsDataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) NumberOfDirectNationalElectionsNelda2015(ctx context.Context, entity string, year int) (*model.NumberOfDirectNationalElectionsNelda2015Dataset, error) {
	var response model.NumberOfDirectNationalElectionsNelda2015Dataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) NumberOfInfantDeathsIhme2017(ctx context.Context, entity string, year int) (*model.NumberOfInfantDeathsIhme2017Dataset, error) {
	var response model.NumberOfInfantDeathsIhme2017Dataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) NumberOfInternetUsersOwidBasedOnWbAndUnwpp(ctx context.Context, entity string, year int) (*model.NumberOfInternetUsersOwidBasedOnWbAndUnwppDataset, error) {
	var response model.NumberOfInternetUsersOwidBasedOnWbAndUnwppDataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) NumberOfNeonatalDeathsIhme2017(ctx context.Context, entity string, year int) (*model.NumberOfNeonatalDeathsIhme2017Dataset, error) {
	var response model.NumberOfNeonatalDeathsIhme2017Dataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) NumberOfObservationsInPovcalPerDecadeOwid2017(ctx context.Context, entity string, year int) (*model.NumberOfObservationsInPovcalPerDecadeOwid2017Dataset, error) {
	var response model.NumberOfObservationsInPovcalPerDecadeOwid2017Dataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) NumberOfPartiesToMultilateralEnvironmentalAgreementsUnctad(ctx context.Context, entity string, year int) (*model.NumberOfPartiesToMultilateralEnvironmentalAgreementsUnctadDataset, error) {
	var response model.NumberOfPartiesToMultilateralEnvironmentalAgreementsUnctadDataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) NumberOfPeopleWhoAreUndernourishedFaoSofi2018AndWorldBank2017(ctx context.Context, entity string, year int) (*model.NumberOfPeopleWhoAreUndernourishedFaoSofi2018AndWorldBank2017Dataset, error) {
	var response model.NumberOfPeopleWhoAreUndernourishedFaoSofi2018AndWorldBank2017Dataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) NumberOfPeopleWithAndWithoutAccessToImprovedSanitationOwidBasedOnWdi(ctx context.Context, entity string, year int) (*model.NumberOfPeopleWithAndWithoutAccessToImprovedSanitationOwidBasedOnWdiDataset, error) {
	var response model.NumberOfPeopleWithAndWithoutAccessToImprovedSanitationOwidBasedOnWdiDataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) NumberOfPeopleWithAndWithoutAccessToImprovedWaterOwidBasedOnWdi(ctx context.Context, entity string, year int) (*model.NumberOfPeopleWithAndWithoutAccessToImprovedWaterOwidBasedOnWdiDataset, error) {
	var response model.NumberOfPeopleWithAndWithoutAccessToImprovedWaterOwidBasedOnWdiDataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) NumberOfPeopleWithAndWithoutAccessToImprovedWaterSourcesWorldBankAndUn(ctx context.Context, entity string, year int) (*model.NumberOfPeopleWithAndWithoutAccessToImprovedWaterSourcesWorldBankAndUnDataset, error) {
	var response model.NumberOfPeopleWithAndWithoutAccessToImprovedWaterSourcesWorldBankAndUnDataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) NumberOfPeopleWithAndWithoutEnergyAccessOwidBasedOnWorldBank2021(ctx context.Context, entity string, year int) (*model.NumberOfPeopleWithAndWithoutEnergyAccessOwidBasedOnWorldBank2021Dataset, error) {
	var response model.NumberOfPeopleWithAndWithoutEnergyAccessOwidBasedOnWorldBank2021Dataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) NumberOfPeopleWithoutAccessToSafeWaterAndSanitationWhoWash2019(ctx context.Context, entity string, year int) (*model.NumberOfPeopleWithoutAccessToSafeWaterAndSanitationWhoWash2019Dataset, error) {
	var response model.NumberOfPeopleWithoutAccessToSafeWaterAndSanitationWhoWash2019Dataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) NumberOfPolioCasesPerOneMillionPopulationWho2017(ctx context.Context, entity string, year int) (*model.NumberOfPolioCasesPerOneMillionPopulationWho2017Dataset, error) {
	var response model.NumberOfPolioCasesPerOneMillionPopulationWho2017Dataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) NumberOfPublishedTitlesSimons2001(ctx context.Context, entity string, year int) (*model.NumberOfPublishedTitlesSimons2001Dataset, error) {
	var response model.NumberOfPublishedTitlesSimons2001Dataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) NumberOfStateBasedConflictsByConflictTypeAndRegionUcdpPrio(ctx context.Context, entity string, year int) (*model.NumberOfStateBasedConflictsByConflictTypeAndRegionUcdpPrioDataset, error) {
	var response model.NumberOfStateBasedConflictsByConflictTypeAndRegionUcdpPrioDataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) NumberOfTouristDeparturesPer1000WorldBankAndUn2019(ctx context.Context, entity string, year int) (*model.NumberOfTouristDeparturesPer1000WorldBankAndUn2019Dataset, error) {
	var response model.NumberOfTouristDeparturesPer1000WorldBankAndUn2019Dataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) OecdTrustInGovernment(ctx context.Context, entity string, year int) (*model.OecdTrustInGovernmentDataset, error) {
	var response model.OecdTrustInGovernmentDataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) OecdEducationPisaTestScoresPisa2015(ctx context.Context, entity string, year int) (*model.OecdEducationPisaTestScoresPisa2015Dataset, error) {
	var response model.OecdEducationPisaTestScoresPisa2015Dataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) OecdSocialSpendingFamily(ctx context.Context, entity string, year int) (*model.OecdSocialSpendingFamilyDataset, error) {
	var response model.OecdSocialSpendingFamilyDataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) OecdSocialSpendingHealth(ctx context.Context, entity string, year int) (*model.OecdSocialSpendingHealthDataset, error) {
	var response model.OecdSocialSpendingHealthDataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) OecdSocialSpendingHousing(ctx context.Context, entity string, year int) (*model.OecdSocialSpendingHousingDataset, error) {
	var response model.OecdSocialSpendingHousingDataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) OecdSocialSpendingIncapacityRelated(ctx context.Context, entity string, year int) (*model.OecdSocialSpendingIncapacityRelatedDataset, error) {
	var response model.OecdSocialSpendingIncapacityRelatedDataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) OecdSocialSpendingOldAge(ctx context.Context, entity string, year int) (*model.OecdSocialSpendingOldAgeDataset, error) {
	var response model.OecdSocialSpendingOldAgeDataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) OecdSocialSpendingOtherSocialPolicyAreas(ctx context.Context, entity string, year int) (*model.OecdSocialSpendingOtherSocialPolicyAreasDataset, error) {
	var response model.OecdSocialSpendingOtherSocialPolicyAreasDataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) OecdSocialSpendingSurvivors(ctx context.Context, entity string, year int) (*model.OecdSocialSpendingSurvivorsDataset, error) {
	var response model.OecdSocialSpendingSurvivorsDataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) OecdSocialSpendingUnemployment(ctx context.Context, entity string, year int) (*model.OecdSocialSpendingUnemploymentDataset, error) {
	var response model.OecdSocialSpendingUnemploymentDataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) OecdConsumptionTaxTrends2016(ctx context.Context, entity string, year int) (*model.OecdConsumptionTaxTrends2016Dataset, error) {
	var response model.OecdConsumptionTaxTrends2016Dataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) OecdEducationStatistics2017(ctx context.Context, entity string, year int) (*model.OecdEducationStatistics2017Dataset, error) {
	var response model.OecdEducationStatistics2017Dataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) OphiMultidimensionalPovertyIndexAlkireAndRobles2016(ctx context.Context, entity string, year int) (*model.OphiMultidimensionalPovertyIndexAlkireAndRobles2016Dataset, error) {
	var response model.OphiMultidimensionalPovertyIndexAlkireAndRobles2016Dataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) OwidCountryToWhoRegions(ctx context.Context, entity string, year int) (*model.OwidCountryToWhoRegionsDataset, error) {
	var response model.OwidCountryToWhoRegionsDataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) OilSpillsItopf2021(ctx context.Context, entity string, year int) (*model.OilSpillsItopf2021Dataset, error) {
	var response model.OilSpillsItopf2021Dataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) OilProductionEtemadAndLuciana(ctx context.Context, entity string, year int) (*model.OilProductionEtemadAndLucianaDataset, error) {
	var response model.OilProductionEtemadAndLucianaDataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) OilandgasEmploymentAndRigCountUsBureauOfLaborStatistics2017(ctx context.Context, entity string, year int) (*model.OilandgasEmploymentAndRigCountUsBureauOfLaborStatistics2017Dataset, error) {
	var response model.OilandgasEmploymentAndRigCountUsBureauOfLaborStatistics2017Dataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) OilcropYieldProductionAndLandUseFao2021(ctx context.Context, entity string, year int) (*model.OilcropYieldProductionAndLandUseFao2021Dataset, error) {
	var response model.OilcropYieldProductionAndLandUseFao2021Dataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) OlympicCompetingNationsAndAthletesOlympicDatabase(ctx context.Context, entity string, year int) (*model.OlympicCompetingNationsAndAthletesOlympicDatabaseDataset, error) {
	var response model.OlympicCompetingNationsAndAthletesOlympicDatabaseDataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) OnshoreWindLcoeIrenaCostDatabase2018(ctx context.Context, entity string, year int) (*model.OnshoreWindLcoeIrenaCostDatabase2018Dataset, error) {
	var response model.OnshoreWindLcoeIrenaCostDatabase2018Dataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) OnshoreWindCostBreakdownIrena2018(ctx context.Context, entity string, year int) (*model.OnshoreWindCostBreakdownIrena2018Dataset, error) {
	var response model.OnshoreWindCostBreakdownIrena2018Dataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) OnshoreWindInstalledProjectCostIrena2018(ctx context.Context, entity string, year int) (*model.OnshoreWindInstalledProjectCostIrena2018Dataset, error) {
	var response model.OnshoreWindInstalledProjectCostIrena2018Dataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) OpioidDeathsDueToOveruseInTheUsCdcWonder2017(ctx context.Context, entity string, year int) (*model.OpioidDeathsDueToOveruseInTheUsCdcWonder2017Dataset, error) {
	var response model.OpioidDeathsDueToOveruseInTheUsCdcWonder2017Dataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) OutputOfKeyIndustrialSectorsInEnglandBankOfEngland2017(ctx context.Context, entity string, year int) (*model.OutputOfKeyIndustrialSectorsInEnglandBankOfEngland2017Dataset, error) {
	var response model.OutputOfKeyIndustrialSectorsInEnglandBankOfEngland2017Dataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) OutputOfKeyIndustriesInEnglandUsingBankOfEngland2017(ctx context.Context, entity string, year int) (*model.OutputOfKeyIndustriesInEnglandUsingBankOfEngland2017Dataset, error) {
	var response model.OutputOfKeyIndustriesInEnglandUsingBankOfEngland2017Dataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) OutputOfKeyServicesSectorsInEnglandUsingBankOfEngland2017(ctx context.Context, entity string, year int) (*model.OutputOfKeyServicesSectorsInEnglandUsingBankOfEngland2017Dataset, error) {
	var response model.OutputOfKeyServicesSectorsInEnglandUsingBankOfEngland2017Dataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) OzoneAndChlorineProjectionsTo2100ScientificAssessment2014(ctx context.Context, entity string, year int) (*model.OzoneAndChlorineProjectionsTo2100ScientificAssessment2014Dataset, error) {
	var response model.OzoneAndChlorineProjectionsTo2100ScientificAssessment2014Dataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) OzoneConcentrationStateofglobalair(ctx context.Context, entity string, year int) (*model.OzoneConcentrationStateofglobalairDataset, error) {
	var response model.OzoneConcentrationStateofglobalairDataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) OzoneDepletingEmissionsSince1960ScientificAssessment2014(ctx context.Context, entity string, year int) (*model.OzoneDepletingEmissionsSince1960ScientificAssessment2014Dataset, error) {
	var response model.OzoneDepletingEmissionsSince1960ScientificAssessment2014Dataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) OzoneDepletionImpactsOnSkinCancerIncidenceSlaperEtAl(ctx context.Context, entity string, year int) (*model.OzoneDepletionImpactsOnSkinCancerIncidenceSlaperEtAlDataset, error) {
	var response model.OzoneDepletionImpactsOnSkinCancerIncidenceSlaperEtAlDataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) OzoneHoleAreaAndConcentrationNasa(ctx context.Context, entity string, year int) (*model.OzoneHoleAreaAndConcentrationNasaDataset, error) {
	var response model.OzoneHoleAreaAndConcentrationNasaDataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) OzoneDepletingEmissionsIndexEea(ctx context.Context, entity string, year int) (*model.OzoneDepletingEmissionsIndexEeaDataset, error) {
	var response model.OzoneDepletingEmissionsIndexEeaDataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) PartiesToMontrealProtocolUnep(ctx context.Context, entity string, year int) (*model.PartiesToMontrealProtocolUnepDataset, error) {
	var response model.PartiesToMontrealProtocolUnepDataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) PatentAndPublicationRatesOwidBasedOnWorldBankAndUn(ctx context.Context, entity string, year int) (*model.PatentAndPublicationRatesOwidBasedOnWorldBankAndUnDataset, error) {
	var response model.PatentAndPublicationRatesOwidBasedOnWorldBankAndUnDataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) PatentsAwardedInEnglandScotlandAndWalesBottomley(ctx context.Context, entity string, year int) (*model.PatentsAwardedInEnglandScotlandAndWalesBottomleyDataset, error) {
	var response model.PatentsAwardedInEnglandScotlandAndWalesBottomleyDataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) PeakFarmlandProjectionAusbuelEtAl2013(ctx context.Context, entity string, year int) (*model.PeakFarmlandProjectionAusbuelEtAl2013Dataset, error) {
	var response model.PeakFarmlandProjectionAusbuelEtAl2013Dataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) PeopleExperiencingHomelessnessInTheUsaPitByShelteringStatusHud2016(ctx context.Context, entity string, year int) (*model.PeopleExperiencingHomelessnessInTheUsaPitByShelteringStatusHud2016Dataset, error) {
	var response model.PeopleExperiencingHomelessnessInTheUsaPitByShelteringStatusHud2016Dataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) PercentageDeathsAttributableToRiskFactorsIhme(ctx context.Context, entity string, year int) (*model.PercentageDeathsAttributableToRiskFactorsIhmeDataset, error) {
	var response model.PercentageDeathsAttributableToRiskFactorsIhmeDataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) PercentageGainedAccessToImprovedWaterAndSanitation1990_2015Who(ctx context.Context, entity string, year int) (*model.PercentageGainedAccessToImprovedWaterAndSanitation19902015WhoDataset, error) {
	var response model.PercentageGainedAccessToImprovedWaterAndSanitation19902015WhoDataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) PercentageOfAmericansLivingAloneByAgeIpums(ctx context.Context, entity string, year int) (*model.PercentageOfAmericansLivingAloneByAgeIpumsDataset, error) {
	var response model.PercentageOfAmericansLivingAloneByAgeIpumsDataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) PercentageOfIndividualsUsingTheInternetIctItu2015(ctx context.Context, entity string, year int) (*model.PercentageOfIndividualsUsingTheInternetIctItu2015Dataset, error) {
	var response model.PercentageOfIndividualsUsingTheInternetIctItu2015Dataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) PercentageOfAdultsLivingAloneInTheUsAndCanadaUsCensusBureauAndStatisticsCanada(ctx context.Context, entity string, year int) (*model.PercentageOfAdultsLivingAloneInTheUsAndCanadaUsCensusBureauAndStatisticsCanadaDataset, error) {
	var response model.PercentageOfAdultsLivingAloneInTheUsAndCanadaUsCensusBureauAndStatisticsCanadaDataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) PercentageOfPersonsWithoutHealthInsuranceCouncilOfEconomicAdvisersAndNationalCenterForHealthStatistics(ctx context.Context, entity string, year int) (*model.PercentageOfPersonsWithoutHealthInsuranceCouncilOfEconomicAdvisersAndNationalCenterForHealthStatisticsDataset, error) {
	var response model.PercentageOfPersonsWithoutHealthInsuranceCouncilOfEconomicAdvisersAndNationalCenterForHealthStatisticsDataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) PerceptionsOfSpendingOnHealthExpenditureIpsos2016(ctx context.Context, entity string, year int) (*model.PerceptionsOfSpendingOnHealthExpenditureIpsos2016Dataset, error) {
	var response model.PerceptionsOfSpendingOnHealthExpenditureIpsos2016Dataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) PhosphateFertilizersFao2017(ctx context.Context, entity string, year int) (*model.PhosphateFertilizersFao2017Dataset, error) {
	var response model.PhosphateFertilizersFao2017Dataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) PiecesOfMailAndNumberOfPostOfficesUnitedStatesPostalService2018(ctx context.Context, entity string, year int) (*model.PiecesOfMailAndNumberOfPostOfficesUnitedStatesPostalService2018Dataset, error) {
	var response model.PiecesOfMailAndNumberOfPostOfficesUnitedStatesPostalService2018Dataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) PlasticWasteJambeckEtAl2015(ctx context.Context, entity string, year int) (*model.PlasticWasteJambeckEtAl2015Dataset, error) {
	var response model.PlasticWasteJambeckEtAl2015Dataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) PlasticBagSubstituteComparisonsDanishEpa2018(ctx context.Context, entity string, year int) (*model.PlasticBagSubstituteComparisonsDanishEpa2018Dataset, error) {
	var response model.PlasticBagSubstituteComparisonsDanishEpa2018Dataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) PlasticDiscardedRecycledIncineratedGeyerEtAl2017(ctx context.Context, entity string, year int) (*model.PlasticDiscardedRecycledIncineratedGeyerEtAl2017Dataset, error) {
	var response model.PlasticDiscardedRecycledIncineratedGeyerEtAl2017Dataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) PlasticImportersToChinaBrooksEtAl2018(ctx context.Context, entity string, year int) (*model.PlasticImportersToChinaBrooksEtAl2018Dataset, error) {
	var response model.PlasticImportersToChinaBrooksEtAl2018Dataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) PlasticImportsByChinaAndImpactOfBanBrooksEtAl2018(ctx context.Context, entity string, year int) (*model.PlasticImportsByChinaAndImpactOfBanBrooksEtAl2018Dataset, error) {
	var response model.PlasticImportsByChinaAndImpactOfBanBrooksEtAl2018Dataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) PlasticOceanPollutionMeijerEtAl2021(ctx context.Context, entity string, year int) (*model.PlasticOceanPollutionMeijerEtAl2021Dataset, error) {
	var response model.PlasticOceanPollutionMeijerEtAl2021Dataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) PlasticPollutionByTop50RiversMeijerEtAl2021(ctx context.Context, entity string, year int) (*model.PlasticPollutionByTop50RiversMeijerEtAl2021Dataset, error) {
	var response model.PlasticPollutionByTop50RiversMeijerEtAl2021Dataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) PlasticProductLifetimeProductionWasteBySourceGeyerEtAl2017(ctx context.Context, entity string, year int) (*model.PlasticProductLifetimeProductionWasteBySourceGeyerEtAl2017Dataset, error) {
	var response model.PlasticProductLifetimeProductionWasteBySourceGeyerEtAl2017Dataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) PlasticWasteGenerationByCountryOwidBasedOnJambeckEtAlAndWorldBank(ctx context.Context, entity string, year int) (*model.PlasticWasteGenerationByCountryOwidBasedOnJambeckEtAlAndWorldBankDataset, error) {
	var response model.PlasticWasteGenerationByCountryOwidBasedOnJambeckEtAlAndWorldBankDataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) PolcalnetGlobalPoverty2017(ctx context.Context, entity string, year int) (*model.PolcalnetGlobalPoverty2017Dataset, error) {
	var response model.PolcalnetGlobalPoverty2017Dataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) PoliticalCompetitionAndParticipationHowWasLifeOecd2014(ctx context.Context, entity string, year int) (*model.PoliticalCompetitionAndParticipationHowWasLifeOecd2014Dataset, error) {
	var response model.PoliticalCompetitionAndParticipationHowWasLifeOecd2014Dataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) PoliticalRegimesBertelsmannTransformationIndex2022(ctx context.Context, entity string, year int) (*model.PoliticalRegimesBertelsmannTransformationIndex2022Dataset, error) {
	var response model.PoliticalRegimesBertelsmannTransformationIndex2022Dataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) PoliticalRegimesEconomistIntelligenceUnit2022(ctx context.Context, entity string, year int) (*model.PoliticalRegimesEconomistIntelligenceUnit2022Dataset, error) {
	var response model.PoliticalRegimesEconomistIntelligenceUnit2022Dataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) PoliticalRegimesFreedomHouse2022(ctx context.Context, entity string, year int) (*model.PoliticalRegimesFreedomHouse2022Dataset, error) {
	var response model.PoliticalRegimesFreedomHouse2022Dataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) PoliticalRegimesOwidBasedOnBoixEtAl2013(ctx context.Context, entity string, year int) (*model.PoliticalRegimesOwidBasedOnBoixEtAl2013Dataset, error) {
	var response model.PoliticalRegimesOwidBasedOnBoixEtAl2013Dataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) PoliticalRegimesOwidBasedOnVDemV12AndLuhrmannEtAl2018(ctx context.Context, entity string, year int) (*model.PoliticalRegimesOwidBasedOnVDemV12AndLuhrmannEtAl2018Dataset, error) {
	var response model.PoliticalRegimesOwidBasedOnVDemV12AndLuhrmannEtAl2018Dataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) PoliticalRegimesPolity5(ctx context.Context, entity string, year int) (*model.PoliticalRegimesPolity5Dataset, error) {
	var response model.PoliticalRegimesPolity5Dataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) PoliticalRegimesSkaaningEtAl2015(ctx context.Context, entity string, year int) (*model.PoliticalRegimesSkaaningEtAl2015Dataset, error) {
	var response model.PoliticalRegimesSkaaningEtAl2015Dataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) PopulationClioInfra2016WithIslandOfIrelandRepNorthern(ctx context.Context, entity string, year int) (*model.PopulationClioInfra2016WithIslandOfIrelandRepNorthernDataset, error) {
	var response model.PopulationClioInfra2016WithIslandOfIrelandRepNorthernDataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) PopulationGapminderHydeAndUn(ctx context.Context, entity string, year int) (*model.PopulationGapminderHydeAndUnDataset, error) {
	var response model.PopulationGapminderHydeAndUnDataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) PopulationDynamicsAndGlobalHumanCapitalIiasa2015(ctx context.Context, entity string, year int) (*model.PopulationDynamicsAndGlobalHumanCapitalIiasa2015Dataset, error) {
	var response model.PopulationDynamicsAndGlobalHumanCapitalIiasa2015Dataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) PopulationEstimatesAndProjectionWittgensteinCentreForDemographyAndGlobalHumanCapital(ctx context.Context, entity string, year int) (*model.PopulationEstimatesAndProjectionWittgensteinCentreForDemographyAndGlobalHumanCapitalDataset, error) {
	var response model.PopulationEstimatesAndProjectionWittgensteinCentreForDemographyAndGlobalHumanCapitalDataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) PopulationGrowthUnPopulationDivision2015Revision(ctx context.Context, entity string, year int) (*model.PopulationGrowthUnPopulationDivision2015RevisionDataset, error) {
	var response model.PopulationGrowthUnPopulationDivision2015RevisionDataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) PopulationByAgeGroupTo2100BasedOnUnwpp2017MediumScenario(ctx context.Context, entity string, year int) (*model.PopulationByAgeGroupTo2100BasedOnUnwpp2017MediumScenarioDataset, error) {
	var response model.PopulationByAgeGroupTo2100BasedOnUnwpp2017MediumScenarioDataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) PopulationByCountry1800To2100GapminderAndUn(ctx context.Context, entity string, year int) (*model.PopulationByCountry1800To2100GapminderAndUnDataset, error) {
	var response model.PopulationByCountry1800To2100GapminderAndUnDataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) PopulationCoveredByTheInternetInternetWorldStats2019(ctx context.Context, entity string, year int) (*model.PopulationCoveredByTheInternetInternetWorldStats2019Dataset, error) {
	var response model.PopulationCoveredByTheInternetInternetWorldStats2019Dataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) PopulationDataGapminderUpTo1949UnPopulationDivision1950To2015(ctx context.Context, entity string, year int) (*model.PopulationDataGapminderUpTo1949UnPopulationDivision1950To2015Dataset, error) {
	var response model.PopulationDataGapminderUpTo1949UnPopulationDivision1950To2015Dataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) PopulationDensityWorldBankGapminderHydeAndUn(ctx context.Context, entity string, year int) (*model.PopulationDensityWorldBankGapminderHydeAndUnDataset, error) {
	var response model.PopulationDensityWorldBankGapminderHydeAndUnDataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) PopulationFedByHaberBoschFertilizersFao2017(ctx context.Context, entity string, year int) (*model.PopulationFedByHaberBoschFertilizersFao2017Dataset, error) {
	var response model.PopulationFedByHaberBoschFertilizersFao2017Dataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) PopulationGrowth1992_2015Listed2017UnPopulationDivision2015(ctx context.Context, entity string, year int) (*model.PopulationGrowth19922015Listed2017UnPopulationDivision2015Dataset, error) {
	var response model.PopulationGrowth19922015Listed2017UnPopulationDivision2015Dataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) PopulationGrowthRateByRegion2020_2100MediumVariantProjectionUnPopulationDivision2015Revision(ctx context.Context, entity string, year int) (*model.PopulationGrowthRateByRegion20202100MediumVariantProjectionUnPopulationDivision2015RevisionDataset, error) {
	var response model.PopulationGrowthRateByRegion20202100MediumVariantProjectionUnPopulationDivision2015RevisionDataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) PopulationUsingInformalSavingPercWorldBankWorldDevelopmentReport2013(ctx context.Context, entity string, year int) (*model.PopulationUsingInformalSavingPercWorldBankWorldDevelopmentReport2013Dataset, error) {
	var response model.PopulationUsingInformalSavingPercWorldBankWorldDevelopmentReport2013Dataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) PostageRatesUnitedStatesPostalService2018(ctx context.Context, entity string, year int) (*model.PostageRatesUnitedStatesPostalService2018Dataset, error) {
	var response model.PostageRatesUnitedStatesPostalService2018Dataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) PotashFertilizersFao2017(ctx context.Context, entity string, year int) (*model.PotashFertilizersFao2017Dataset, error) {
	var response model.PotashFertilizersFao2017Dataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) PovertyHeadcountAtMoney190ADay2011PppHighIncomeWorldBankPovcal2017(ctx context.Context, entity string, year int) (*model.PovertyHeadcountAtMoney190ADay2011PppHighIncomeWorldBankPovcal2017Dataset, error) {
	var response model.PovertyHeadcountAtMoney190ADay2011PppHighIncomeWorldBankPovcal2017Dataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) PovertyRateLess50percOfMedianLisKeyFigures2018(ctx context.Context, entity string, year int) (*model.PovertyRateLess50percOfMedianLisKeyFigures2018Dataset, error) {
	var response model.PovertyRateLess50percOfMedianLisKeyFigures2018Dataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) PrecipitationAnomalyInUsNoaa(ctx context.Context, entity string, year int) (*model.PrecipitationAnomalyInUsNoaaDataset, error) {
	var response model.PrecipitationAnomalyInUsNoaaDataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) PressFreedomFreedomHouse2017(ctx context.Context, entity string, year int) (*model.PressFreedomFreedomHouse2017Dataset, error) {
	var response model.PressFreedomFreedomHouse2017Dataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) PrevalenceOfUndernourishmentInDevelopingCountriesFaoFoodSecurityIndicators2017(ctx context.Context, entity string, year int) (*model.PrevalenceOfUndernourishmentInDevelopingCountriesFaoFoodSecurityIndicators2017Dataset, error) {
	var response model.PrevalenceOfUndernourishmentInDevelopingCountriesFaoFoodSecurityIndicators2017Dataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) PrevalenceOfVitaminADeficiencyInChildrenWho2017(ctx context.Context, entity string, year int) (*model.PrevalenceOfVitaminADeficiencyInChildrenWho2017Dataset, error) {
	var response model.PrevalenceOfVitaminADeficiencyInChildrenWho2017Dataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) PrevalenceOfAlcoholDrinkingInTheUsaCdc(ctx context.Context, entity string, year int) (*model.PrevalenceOfAlcoholDrinkingInTheUsaCdcDataset, error) {
	var response model.PrevalenceOfAlcoholDrinkingInTheUsaCdcDataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) PrevalenceOfUndernourishmentWorldBank2017AndUnSofi2018(ctx context.Context, entity string, year int) (*model.PrevalenceOfUndernourishmentWorldBank2017AndUnSofi2018Dataset, error) {
	var response model.PrevalenceOfUndernourishmentWorldBank2017AndUnSofi2018Dataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) PrevalenceOfUndernourishmentByRegionUnFaoSofi2017And2018(ctx context.Context, entity string, year int) (*model.PrevalenceOfUndernourishmentByRegionUnFaoSofi2017And2018Dataset, error) {
	var response model.PrevalenceOfUndernourishmentByRegionUnFaoSofi2017And2018Dataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) PrevalenceOfUndernourishmentSince2000Faostats2018(ctx context.Context, entity string, year int) (*model.PrevalenceOfUndernourishmentSince2000Faostats2018Dataset, error) {
	var response model.PrevalenceOfUndernourishmentSince2000Faostats2018Dataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) PrevalenceOfVitaminADeficiencyInPregnantWomenWho2009(ctx context.Context, entity string, year int) (*model.PrevalenceOfVitaminADeficiencyInPregnantWomenWho2009Dataset, error) {
	var response model.PrevalenceOfVitaminADeficiencyInPregnantWomenWho2009Dataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) PrevalenceOfWeightCategoriesInFemalesNcdrisc2017(ctx context.Context, entity string, year int) (*model.PrevalenceOfWeightCategoriesInFemalesNcdrisc2017Dataset, error) {
	var response model.PrevalenceOfWeightCategoriesInFemalesNcdrisc2017Dataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) PrevalenceOfWeightCategoriesInMalesNcdrisc2017(ctx context.Context, entity string, year int) (*model.PrevalenceOfWeightCategoriesInMalesNcdrisc2017Dataset, error) {
	var response model.PrevalenceOfWeightCategoriesInMalesNcdrisc2017Dataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) PrevalenceOfZincDeficiencyWessellsEtAl2012(ctx context.Context, entity string, year int) (*model.PrevalenceOfZincDeficiencyWessellsEtAl2012Dataset, error) {
	var response model.PrevalenceOfZincDeficiencyWessellsEtAl2012Dataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) PriceForLightFouquet(ctx context.Context, entity string, year int) (*model.PriceForLightFouquetDataset, error) {
	var response model.PriceForLightFouquetDataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) PriceOfMobileDataAllianceForAffordableInternet2019(ctx context.Context, entity string, year int) (*model.PriceOfMobileDataAllianceForAffordableInternet2019Dataset, error) {
	var response model.PriceOfMobileDataAllianceForAffordableInternet2019Dataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) PriceOfNailsSince1695DanielSichels2017(ctx context.Context, entity string, year int) (*model.PriceOfNailsSince1695DanielSichels2017Dataset, error) {
	var response model.PriceOfNailsSince1695DanielSichels2017Dataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) PrimaryEnergyConsumptionBpAndShift2020(ctx context.Context, entity string, year int) (*model.PrimaryEnergyConsumptionBpAndShift2020Dataset, error) {
	var response model.PrimaryEnergyConsumptionBpAndShift2020Dataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) PrimaryEnergyConsumptionBpAndEia2022(ctx context.Context, entity string, year int) (*model.PrimaryEnergyConsumptionBpAndEia2022Dataset, error) {
	var response model.PrimaryEnergyConsumptionBpAndEia2022Dataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) PrisonersPer100_000FromWorldPrisonBriefDownloadedSeptember2018CountryStandardized(ctx context.Context, entity string, year int) (*model.PrisonersPer100000FromWorldPrisonBriefDownloadedSeptember2018CountryStandardizedDataset, error) {
	var response model.PrisonersPer100000FromWorldPrisonBriefDownloadedSeptember2018CountryStandardizedDataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) ProjectedChangeInUnder5PopulationByCountry2015_2050_2100OwidBasedOnUnPopulation2017(ctx context.Context, entity string, year int) (*model.ProjectedChangeInUnder5PopulationByCountry201520502100OwidBasedOnUnPopulation2017Dataset, error) {
	var response model.ProjectedChangeInUnder5PopulationByCountry201520502100OwidBasedOnUnPopulation2017Dataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) ProjectedExtremePovertyAmongDifferentGroupingsOfFragileStatesCrespoCuaresmaEtAl2018OecdWorldBank(ctx context.Context, entity string, year int) (*model.ProjectedExtremePovertyAmongDifferentGroupingsOfFragileStatesCrespoCuaresmaEtAl2018OecdWorldBankDataset, error) {
	var response model.ProjectedExtremePovertyAmongDifferentGroupingsOfFragileStatesCrespoCuaresmaEtAl2018OecdWorldBankDataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) ProjectionsOfPeakAgriculturalLandFao2006Oecd2012Mea2005(ctx context.Context, entity string, year int) (*model.ProjectionsOfPeakAgriculturalLandFao2006Oecd2012Mea2005Dataset, error) {
	var response model.ProjectionsOfPeakAgriculturalLandFao2006Oecd2012Mea2005Dataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) ProportionOf45_54YearOldsWithTertiaryEducation2009Oecd2012(ctx context.Context, entity string, year int) (*model.ProportionOf4554YearOldsWithTertiaryEducation2009Oecd2012Dataset, error) {
	var response model.ProportionOf4554YearOldsWithTertiaryEducation2009Oecd2012Dataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) PublicExpenditureOnEducationOecdTanziAndSchuknecht2000(ctx context.Context, entity string, year int) (*model.PublicExpenditureOnEducationOecdTanziAndSchuknecht2000Dataset, error) {
	var response model.PublicExpenditureOnEducationOecdTanziAndSchuknecht2000Dataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) PublicSupportAndOppositionToNuclearEnergyIpsosMori2011(ctx context.Context, entity string, year int) (*model.PublicSupportAndOppositionToNuclearEnergyIpsosMori2011Dataset, error) {
	var response model.PublicSupportAndOppositionToNuclearEnergyIpsosMori2011Dataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) RandDatabaseOfWorldwideTerrorismIncidents(ctx context.Context, entity string, year int) (*model.RandDatabaseOfWorldwideTerrorismIncidentsDataset, error) {
	var response model.RandDatabaseOfWorldwideTerrorismIncidentsDataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) RaisedBloodPressurePrevalenceNcdRisc2017(ctx context.Context, entity string, year int) (*model.RaisedBloodPressurePrevalenceNcdRisc2017Dataset, error) {
	var response model.RaisedBloodPressurePrevalenceNcdRisc2017Dataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) RateOfNaturalPopulationIncreaseUnPopulationDivision2015(ctx context.Context, entity string, year int) (*model.RateOfNaturalPopulationIncreaseUnPopulationDivision2015Dataset, error) {
	var response model.RateOfNaturalPopulationIncreaseUnPopulationDivision2015Dataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) RateOfInternationallyObservedElectionsHydeAndMarinov2012(ctx context.Context, entity string, year int) (*model.RateOfInternationallyObservedElectionsHydeAndMarinov2012Dataset, error) {
	var response model.RateOfInternationallyObservedElectionsHydeAndMarinov2012Dataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) RateOfInternationallyObservedElectionsSusanHyde2011(ctx context.Context, entity string, year int) (*model.RateOfInternationallyObservedElectionsSusanHyde2011Dataset, error) {
	var response model.RateOfInternationallyObservedElectionsSusanHyde2011Dataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) RealCommodityPriceIndexSince1850Jacks2016(ctx context.Context, entity string, year int) (*model.RealCommodityPriceIndexSince1850Jacks2016Dataset, error) {
	var response model.RealCommodityPriceIndexSince1850Jacks2016Dataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) RealGdpPerCapitaLondonAndDelhiOwid(ctx context.Context, entity string, year int) (*model.RealGdpPerCapitaLondonAndDelhiOwidDataset, error) {
	var response model.RealGdpPerCapitaLondonAndDelhiOwidDataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) RecycledPlasticExportsBrooksEtAl2018(ctx context.Context, entity string, year int) (*model.RecycledPlasticExportsBrooksEtAl2018Dataset, error) {
	var response model.RecycledPlasticExportsBrooksEtAl2018Dataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) RegimePopulationsOwidBasedOnLuhrmannEtAl2018VDemV12Owid2021GapminderV6HydeV32AndUn2019(ctx context.Context, entity string, year int) (*model.RegimePopulationsOwidBasedOnLuhrmannEtAl2018VDemV12Owid2021GapminderV6HydeV32AndUn2019Dataset, error) {
	var response model.RegimePopulationsOwidBasedOnLuhrmannEtAl2018VDemV12Owid2021GapminderV6HydeV32AndUn2019Dataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) RelativeEarningsOfAdultsByEducationalAttainmentEducationAtAGlance2017OecdIndicators2017(ctx context.Context, entity string, year int) (*model.RelativeEarningsOfAdultsByEducationalAttainmentEducationAtAGlance2017OecdIndicators2017Dataset, error) {
	var response model.RelativeEarningsOfAdultsByEducationalAttainmentEducationAtAGlance2017OecdIndicators2017Dataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) RelativeWagesOfCraftsmenToLabourers1200_2000Clark2005(ctx context.Context, entity string, year int) (*model.RelativeWagesOfCraftsmenToLabourers12002000Clark2005Dataset, error) {
	var response model.RelativeWagesOfCraftsmenToLabourers12002000Clark2005Dataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) RenewableEnergyPercElectricityProductionWorldBank2015(ctx context.Context, entity string, year int) (*model.RenewableEnergyPercElectricityProductionWorldBank2015Dataset, error) {
	var response model.RenewableEnergyPercElectricityProductionWorldBank2015Dataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) RenewableEnergyCapacityByRegionIrena2017(ctx context.Context, entity string, year int) (*model.RenewableEnergyCapacityByRegionIrena2017Dataset, error) {
	var response model.RenewableEnergyCapacityByRegionIrena2017Dataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) RenewableEnergyCapacityByTechnologyIrena2017(ctx context.Context, entity string, year int) (*model.RenewableEnergyCapacityByTechnologyIrena2017Dataset, error) {
	var response model.RenewableEnergyCapacityByTechnologyIrena2017Dataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) RenewableInvestmentAsPercOfGdpBnepAndWorldBank(ctx context.Context, entity string, year int) (*model.RenewableInvestmentAsPercOfGdpBnepAndWorldBankDataset, error) {
	var response model.RenewableInvestmentAsPercOfGdpBnepAndWorldBankDataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) RenewableEnergyCostsIrena2020(ctx context.Context, entity string, year int) (*model.RenewableEnergyCostsIrena2020Dataset, error) {
	var response model.RenewableEnergyCostsIrena2020Dataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) RenewablesPatentsIrena2016(ctx context.Context, entity string, year int) (*model.RenewablesPatentsIrena2016Dataset, error) {
	var response model.RenewablesPatentsIrena2016Dataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) ReportedGuineaWormCasesWho2021(ctx context.Context, entity string, year int) (*model.ReportedGuineaWormCasesWho2021Dataset, error) {
	var response model.ReportedGuineaWormCasesWho2021Dataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) ReportedNumberAndDifferentEstimationsOfPolioCasesWho2018(ctx context.Context, entity string, year int) (*model.ReportedNumberAndDifferentEstimationsOfPolioCasesWho2018Dataset, error) {
	var response model.ReportedNumberAndDifferentEstimationsOfPolioCasesWho2018Dataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) RequiredRateOfMaternalMortalityDeclineForSdgBasedOnWorldBank2018(ctx context.Context, entity string, year int) (*model.RequiredRateOfMaternalMortalityDeclineForSdgBasedOnWorldBank2018Dataset, error) {
	var response model.RequiredRateOfMaternalMortalityDeclineForSdgBasedOnWorldBank2018Dataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) ReservesProductionRatioBpStatistics2016(ctx context.Context, entity string, year int) (*model.ReservesProductionRatioBpStatistics2016Dataset, error) {
	var response model.ReservesProductionRatioBpStatistics2016Dataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) RevenueSharesFromTaxFlora1983AndIctd2016(ctx context.Context, entity string, year int) (*model.RevenueSharesFromTaxFlora1983AndIctd2016Dataset, error) {
	var response model.RevenueSharesFromTaxFlora1983AndIctd2016Dataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) RhinoPoachingRatesAfrsg2019(ctx context.Context, entity string, year int) (*model.RhinoPoachingRatesAfrsg2019Dataset, error) {
	var response model.RhinoPoachingRatesAfrsg2019Dataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) RhinoPopulationsAfrsgAndOtherSources2022(ctx context.Context, entity string, year int) (*model.RhinoPopulationsAfrsgAndOtherSources2022Dataset, error) {
	var response model.RhinoPopulationsAfrsgAndOtherSources2022Dataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) RiskAttributionOfCancerDeathsToTobaccoSmokingIhme(ctx context.Context, entity string, year int) (*model.RiskAttributionOfCancerDeathsToTobaccoSmokingIhmeDataset, error) {
	var response model.RiskAttributionOfCancerDeathsToTobaccoSmokingIhmeDataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) RoadDeathsAndInjuriesOecd(ctx context.Context, entity string, year int) (*model.RoadDeathsAndInjuriesOecdDataset, error) {
	var response model.RoadDeathsAndInjuriesOecdDataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) RotavirusDeathsAndCasesInUnder5sIhme2018(ctx context.Context, entity string, year int) (*model.RotavirusDeathsAndCasesInUnder5sIhme2018Dataset, error) {
	var response model.RotavirusDeathsAndCasesInUnder5sIhme2018Dataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) RoughSleepingInEnglandInThe2010sOwidBasedOnUkNationalStatistics2018(ctx context.Context, entity string, year int) (*model.RoughSleepingInEnglandInThe2010sOwidBasedOnUkNationalStatistics2018Dataset, error) {
	var response model.RoughSleepingInEnglandInThe2010sOwidBasedOnUkNationalStatistics2018Dataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) SipriMilitaryExpenditureDatabase(ctx context.Context, entity string, year int) (*model.SipriMilitaryExpenditureDatabaseDataset, error) {
	var response model.SipriMilitaryExpenditureDatabaseDataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) So2EmissionsByCountry1850_2000ClioInfra(ctx context.Context, entity string, year int) (*model.So2EmissionsByCountry18502000ClioInfraDataset, error) {
	var response model.So2EmissionsByCountry18502000ClioInfraDataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) So2EmissionsByRegionOecd2014AndKlimontEtAl2013(ctx context.Context, entity string, year int) (*model.So2EmissionsByRegionOecd2014AndKlimontEtAl2013Dataset, error) {
	var response model.So2EmissionsByRegionOecd2014AndKlimontEtAl2013Dataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) So2EmissionsChinaAndIndiaKlimontEtAl2013(ctx context.Context, entity string, year int) (*model.So2EmissionsChinaAndIndiaKlimontEtAl2013Dataset, error) {
	var response model.So2EmissionsChinaAndIndiaKlimontEtAl2013Dataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) So2PerCapitaClioInfra(ctx context.Context, entity string, year int) (*model.So2PerCapitaClioInfraDataset, error) {
	var response model.So2PerCapitaClioInfraDataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) SameSexMarriageLawPewResearchCenterCfr(ctx context.Context, entity string, year int) (*model.SameSexMarriageLawPewResearchCenterCfrDataset, error) {
	var response model.SameSexMarriageLawPewResearchCenterCfrDataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) SameSexMarriagesBySexInTheNetherlandsCbs2016(ctx context.Context, entity string, year int) (*model.SameSexMarriagesBySexInTheNetherlandsCbs2016Dataset, error) {
	var response model.SameSexMarriagesBySexInTheNetherlandsCbs2016Dataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) SameSexMarriedHouseholdsInTheUs(ctx context.Context, entity string, year int) (*model.SameSexMarriedHouseholdsInTheUsDataset, error) {
	var response model.SameSexMarriedHouseholdsInTheUsDataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) SelfReportedLonelinessInOlderAdultsOwid2018(ctx context.Context, entity string, year int) (*model.SelfReportedLonelinessInOlderAdultsOwid2018Dataset, error) {
	var response model.SelfReportedLonelinessInOlderAdultsOwid2018Dataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) SexRatioAtBirthChaoEtAl2019(ctx context.Context, entity string, year int) (*model.SexRatioAtBirthChaoEtAl2019Dataset, error) {
	var response model.SexRatioAtBirthChaoEtAl2019Dataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) SexRatioAtBirthByBirthOrderInSkoreaAndChinaJiangEtAl2017AndNsoKorea(ctx context.Context, entity string, year int) (*model.SexRatioAtBirthByBirthOrderInSkoreaAndChinaJiangEtAl2017AndNsoKoreaDataset, error) {
	var response model.SexRatioAtBirthByBirthOrderInSkoreaAndChinaJiangEtAl2017AndNsoKoreaDataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) SexRatioByAgeOwidBasedOnUnwpp2017(ctx context.Context, entity string, year int) (*model.SexRatioByAgeOwidBasedOnUnwpp2017Dataset, error) {
	var response model.SexRatioByAgeOwidBasedOnUnwpp2017Dataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) SexualViolenceUnicef2017(ctx context.Context, entity string, year int) (*model.SexualViolenceUnicef2017Dataset, error) {
	var response model.SexualViolenceUnicef2017Dataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) ShareOfTop1percInNetPersonalWealthWorldWealthAndIncomeDatabase2018(ctx context.Context, entity string, year int) (*model.ShareOfTop1percInNetPersonalWealthWorldWealthAndIncomeDatabase2018Dataset, error) {
	var response model.ShareOfTop1percInNetPersonalWealthWorldWealthAndIncomeDatabase2018Dataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) ShareOfArableLandWhichIsOrganicOwidBasedOnFao(ctx context.Context, entity string, year int) (*model.ShareOfArableLandWhichIsOrganicOwidBasedOnFaoDataset, error) {
	var response model.ShareOfArableLandWhichIsOrganicOwidBasedOnFaoDataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) ShareOfCountriesWhereHomosexualityIsLegalOwidBasedOnKennyAndPatel2017(ctx context.Context, entity string, year int) (*model.ShareOfCountriesWhereHomosexualityIsLegalOwidBasedOnKennyAndPatel2017Dataset, error) {
	var response model.ShareOfCountriesWhereHomosexualityIsLegalOwidBasedOnKennyAndPatel2017Dataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) ShareOfDeathsAttributedToAirPollutionIhme2019(ctx context.Context, entity string, year int) (*model.ShareOfDeathsAttributedToAirPollutionIhme2019Dataset, error) {
	var response model.ShareOfDeathsAttributedToAirPollutionIhme2019Dataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) ShareOfEmploymentInTheFinancialSectorGgdc2017(ctx context.Context, entity string, year int) (*model.ShareOfEmploymentInTheFinancialSectorGgdc2017Dataset, error) {
	var response model.ShareOfEmploymentInTheFinancialSectorGgdc2017Dataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) ShareOfEnergyFromCerealsRootsAndTubersFao2017(ctx context.Context, entity string, year int) (*model.ShareOfEnergyFromCerealsRootsAndTubersFao2017Dataset, error) {
	var response model.ShareOfEnergyFromCerealsRootsAndTubersFao2017Dataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) ShareOfFoodLostByFoodTypeAndRegionFao2019(ctx context.Context, entity string, year int) (*model.ShareOfFoodLostByFoodTypeAndRegionFao2019Dataset, error) {
	var response model.ShareOfFoodLostByFoodTypeAndRegionFao2019Dataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) ShareOfLandownersWhoAreFemaleFao2017(ctx context.Context, entity string, year int) (*model.ShareOfLandownersWhoAreFemaleFao2017Dataset, error) {
	var response model.ShareOfLandownersWhoAreFemaleFao2017Dataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) ShareOfMarriagesInEnglandAndWalesThatEndedInDivorceUkOns2020(ctx context.Context, entity string, year int) (*model.ShareOfMarriagesInEnglandAndWalesThatEndedInDivorceUkOns2020Dataset, error) {
	var response model.ShareOfMarriagesInEnglandAndWalesThatEndedInDivorceUkOns2020Dataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) ShareOfPeopleExperiencingHomelessnessInTheUsa2007_2016Per100_000Hud2016AndUsCensusBureau2010(ctx context.Context, entity string, year int) (*model.ShareOfPeopleExperiencingHomelessnessInTheUsa20072016Per100000Hud2016AndUsCensusBureau2010Dataset, error) {
	var response model.ShareOfPeopleExperiencingHomelessnessInTheUsa20072016Per100000Hud2016AndUsCensusBureau2010Dataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) ShareOfPeopleWhoReportHavingIntentionsToStartBusinessGlobalEntrepreneurshipMonitor(ctx context.Context, entity string, year int) (*model.ShareOfPeopleWhoReportHavingIntentionsToStartBusinessGlobalEntrepreneurshipMonitorDataset, error) {
	var response model.ShareOfPeopleWhoReportHavingIntentionsToStartBusinessGlobalEntrepreneurshipMonitorDataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) ShareOfPopulationCoveredBySocialProtectionAspireWorldBank2019(ctx context.Context, entity string, year int) (*model.ShareOfPopulationCoveredBySocialProtectionAspireWorldBank2019Dataset, error) {
	var response model.ShareOfPopulationCoveredBySocialProtectionAspireWorldBank2019Dataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) ShareOfPrimarySchoolChildrenAchievingMinimumReadingProficiencyRichVsPoorUnesco(ctx context.Context, entity string, year int) (*model.ShareOfPrimarySchoolChildrenAchievingMinimumReadingProficiencyRichVsPoorUnescoDataset, error) {
	var response model.ShareOfPrimarySchoolChildrenAchievingMinimumReadingProficiencyRichVsPoorUnescoDataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) ShareOfServicesInTotalExportsWdi2017(ctx context.Context, entity string, year int) (*model.ShareOfServicesInTotalExportsWdi2017Dataset, error) {
	var response model.ShareOfServicesInTotalExportsWdi2017Dataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) ShareOfSingleParentFamiliesUnPopulationDivision2018(ctx context.Context, entity string, year int) (*model.ShareOfSingleParentFamiliesUnPopulationDivision2018Dataset, error) {
	var response model.ShareOfSingleParentFamiliesUnPopulationDivision2018Dataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) ShareOfWomenInTopIncomeGroupsAtkinsonCasaricoAndVoitchovsky2018Old(ctx context.Context, entity string, year int) (*model.ShareOfWomenInTopIncomeGroupsAtkinsonCasaricoAndVoitchovsky2018OldDataset, error) {
	var response model.ShareOfWomenInTopIncomeGroupsAtkinsonCasaricoAndVoitchovsky2018OldDataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) ShareOfWorldMerchandiseTradeByTypeOfTradeFouquinAndHugotCepii2016DyadicData(ctx context.Context, entity string, year int) (*model.ShareOfWorldMerchandiseTradeByTypeOfTradeFouquinAndHugotCepii2016DyadicDataDataset, error) {
	var response model.ShareOfWorldMerchandiseTradeByTypeOfTradeFouquinAndHugotCepii2016DyadicDataDataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) SharkAttacksAndFatalitiesGlobalSharkAttackFileGsaf2018(ctx context.Context, entity string, year int) (*model.SharkAttacksAndFatalitiesGlobalSharkAttackFileGsaf2018Dataset, error) {
	var response model.SharkAttacksAndFatalitiesGlobalSharkAttackFileGsaf2018Dataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) SignificantEarthquakeEventsNgdcNasa(ctx context.Context, entity string, year int) (*model.SignificantEarthquakeEventsNgdcNasaDataset, error) {
	var response model.SignificantEarthquakeEventsNgdcNasaDataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) SignificantVolcanicEruptionsNgdcWds(ctx context.Context, entity string, year int) (*model.SignificantVolcanicEruptionsNgdcWdsDataset, error) {
	var response model.SignificantVolcanicEruptionsNgdcWdsDataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) SmallpoxCasesByCountry1920_1977(ctx context.Context, entity string, year int) (*model.SmallpoxCasesByCountry19201977Dataset, error) {
	var response model.SmallpoxCasesByCountry19201977Dataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) SmallpoxCasesReportedAndRevisedFennerEtAl1988(ctx context.Context, entity string, year int) (*model.SmallpoxCasesReportedAndRevisedFennerEtAl1988Dataset, error) {
	var response model.SmallpoxCasesReportedAndRevisedFennerEtAl1988Dataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) SmokingCigaretteSalesInternationalSmokingStatistics2017(ctx context.Context, entity string, year int) (*model.SmokingCigaretteSalesInternationalSmokingStatistics2017Dataset, error) {
	var response model.SmokingCigaretteSalesInternationalSmokingStatistics2017Dataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) SmokingPrevalenceAndCigaretteConsumptionIhmeGhdx2012(ctx context.Context, entity string, year int) (*model.SmokingPrevalenceAndCigaretteConsumptionIhmeGhdx2012Dataset, error) {
	var response model.SmokingPrevalenceAndCigaretteConsumptionIhmeGhdx2012Dataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) SocialExpenditureInTheLongRunLindert2004Oecd1985OecdSocx(ctx context.Context, entity string, year int) (*model.SocialExpenditureInTheLongRunLindert2004Oecd1985OecdSocxDataset, error) {
	var response model.SocialExpenditureInTheLongRunLindert2004Oecd1985OecdSocxDataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) SolarPvModuleCostsAndCapacityLafondEtAl2017AndIrena(ctx context.Context, entity string, year int) (*model.SolarPvModuleCostsAndCapacityLafondEtAl2017AndIrenaDataset, error) {
	var response model.SolarPvModuleCostsAndCapacityLafondEtAl2017AndIrenaDataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) SolarPvSystemsCostsBarboseAndDarghouth2016(ctx context.Context, entity string, year int) (*model.SolarPvSystemsCostsBarboseAndDarghouth2016Dataset, error) {
	var response model.SolarPvSystemsCostsBarboseAndDarghouth2016Dataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) SolidFuelUseForCookingByRegionBonjourEtAl2013(ctx context.Context, entity string, year int) (*model.SolidFuelUseForCookingByRegionBonjourEtAl2013Dataset, error) {
	var response model.SolidFuelUseForCookingByRegionBonjourEtAl2013Dataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) SplitOfExportsToDifferentCountryGroupsOwidCalculationsBasedOnFouquinAndHugotCepii2016DyadicData(ctx context.Context, entity string, year int) (*model.SplitOfExportsToDifferentCountryGroupsOwidCalculationsBasedOnFouquinAndHugotCepii2016DyadicDataDataset, error) {
	var response model.SplitOfExportsToDifferentCountryGroupsOwidCalculationsBasedOnFouquinAndHugotCepii2016DyadicDataDataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) StateOfVaccineConfidenceLarsonEtAl2016(ctx context.Context, entity string, year int) (*model.StateOfVaccineConfidenceLarsonEtAl2016Dataset, error) {
	var response model.StateOfVaccineConfidenceLarsonEtAl2016Dataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) StateBasedConflictDeathsSince1946ByRegionAndConflictTypePrioUcdp2022(ctx context.Context, entity string, year int) (*model.StateBasedConflictDeathsSince1946ByRegionAndConflictTypePrioUcdp2022Dataset, error) {
	var response model.StateBasedConflictDeathsSince1946ByRegionAndConflictTypePrioUcdp2022Dataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) SubnationalInequalityOecdBasedOnRoyuelaEtAl2014(ctx context.Context, entity string, year int) (*model.SubnationalInequalityOecdBasedOnRoyuelaEtAl2014Dataset, error) {
	var response model.SubnationalInequalityOecdBasedOnRoyuelaEtAl2014Dataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) SuicideRatesBySexAndAgeIhme2019(ctx context.Context, entity string, year int) (*model.SuicideRatesBySexAndAgeIhme2019Dataset, error) {
	var response model.SuicideRatesBySexAndAgeIhme2019Dataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) SuicidesFromPesticidesMewEtAl2017(ctx context.Context, entity string, year int) (*model.SuicidesFromPesticidesMewEtAl2017Dataset, error) {
	var response model.SuicidesFromPesticidesMewEtAl2017Dataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) SupercomputerPowerFlopsTop500Database(ctx context.Context, entity string, year int) (*model.SupercomputerPowerFlopsTop500DatabaseDataset, error) {
	var response model.SupercomputerPowerFlopsTop500DatabaseDataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) SurfaceOceanPlasticByMassEriksenEtAl2014(ctx context.Context, entity string, year int) (*model.SurfaceOceanPlasticByMassEriksenEtAl2014Dataset, error) {
	var response model.SurfaceOceanPlasticByMassEriksenEtAl2014Dataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) SurfaceOceanPlasticByParticleCountEriksenEtAl2014(ctx context.Context, entity string, year int) (*model.SurfaceOceanPlasticByParticleCountEriksenEtAl2014Dataset, error) {
	var response model.SurfaceOceanPlasticByParticleCountEriksenEtAl2014Dataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) SwedishHistoricalNationalAccountsSchonAndKrantz2007_2012_2015(ctx context.Context, entity string, year int) (*model.SwedishHistoricalNationalAccountsSchonAndKrantz200720122015Dataset, error) {
	var response model.SwedishHistoricalNationalAccountsSchonAndKrantz200720122015Dataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) TaxRevenuePiketty2014(ctx context.Context, entity string, year int) (*model.TaxRevenuePiketty2014Dataset, error) {
	var response model.TaxRevenuePiketty2014Dataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) TaxCompositionArroyoAbadAndPLindert2016(ctx context.Context, entity string, year int) (*model.TaxCompositionArroyoAbadAndPLindert2016Dataset, error) {
	var response model.TaxCompositionArroyoAbadAndPLindert2016Dataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) TaxCompositionTodaroAndSmith2014(ctx context.Context, entity string, year int) (*model.TaxCompositionTodaroAndSmith2014Dataset, error) {
	var response model.TaxCompositionTodaroAndSmith2014Dataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) TaxesIctdGrd2021(ctx context.Context, entity string, year int) (*model.TaxesIctdGrd2021Dataset, error) {
	var response model.TaxesIctdGrd2021Dataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) TeacherAbsenteeismBoldEtAl2017(ctx context.Context, entity string, year int) (*model.TeacherAbsenteeismBoldEtAl2017Dataset, error) {
	var response model.TeacherAbsenteeismBoldEtAl2017Dataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) TeachingTimeLostWorldDevelopmentReport2018(ctx context.Context, entity string, year int) (*model.TeachingTimeLostWorldDevelopmentReport2018Dataset, error) {
	var response model.TeachingTimeLostWorldDevelopmentReport2018Dataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) TechnologyAdoptionIsard1942AndOthers(ctx context.Context, entity string, year int) (*model.TechnologyAdoptionIsard1942AndOthersDataset, error) {
	var response model.TechnologyAdoptionIsard1942AndOthersDataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) TechnologyDiffusionCominAndHobijn2004AndOthers(ctx context.Context, entity string, year int) (*model.TechnologyDiffusionCominAndHobijn2004AndOthersDataset, error) {
	var response model.TechnologyDiffusionCominAndHobijn2004AndOthersDataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) TemporaryAccommodationInEnglandUkGovernment2018(ctx context.Context, entity string, year int) (*model.TemporaryAccommodationInEnglandUkGovernment2018Dataset, error) {
	var response model.TemporaryAccommodationInEnglandUkGovernment2018Dataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) TerrainRuggednessIndexNunnAndPuga2012(ctx context.Context, entity string, year int) (*model.TerrainRuggednessIndexNunnAndPuga2012Dataset, error) {
	var response model.TerrainRuggednessIndexNunnAndPuga2012Dataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) TerrorismIncidentsFatalitiesAndInjuriesGlobalTerrorismDatabase2018(ctx context.Context, entity string, year int) (*model.TerrorismIncidentsFatalitiesAndInjuriesGlobalTerrorismDatabase2018Dataset, error) {
	var response model.TerrorismIncidentsFatalitiesAndInjuriesGlobalTerrorismDatabase2018Dataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) TerroristAttackByTargetTypeGlobalTerrorismDatabase2018(ctx context.Context, entity string, year int) (*model.TerroristAttackByTargetTypeGlobalTerrorismDatabase2018Dataset, error) {
	var response model.TerroristAttackByTargetTypeGlobalTerrorismDatabase2018Dataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) TerroristAttacksByTypeGlobalTerrorismDatabase2018(ctx context.Context, entity string, year int) (*model.TerroristAttacksByTypeGlobalTerrorismDatabase2018Dataset, error) {
	var response model.TerroristAttacksByTypeGlobalTerrorismDatabase2018Dataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) TerroristAttacksByWeaponTypeGlobalTerrorismDatabase2018(ctx context.Context, entity string, year int) (*model.TerroristAttacksByWeaponTypeGlobalTerrorismDatabase2018Dataset, error) {
	var response model.TerroristAttacksByWeaponTypeGlobalTerrorismDatabase2018Dataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) TetanusNeonatalRateCalculatedFromWhoIncidence2017AndWdiPopulationDataHannahBehrens(ctx context.Context, entity string, year int) (*model.TetanusNeonatalRateCalculatedFromWhoIncidence2017AndWdiPopulationDataHannahBehrensDataset, error) {
	var response model.TetanusNeonatalRateCalculatedFromWhoIncidence2017AndWdiPopulationDataHannahBehrensDataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) TheAllocationOfTimeOverFiveDecadesAguiarAndHurst2006(ctx context.Context, entity string, year int) (*model.TheAllocationOfTimeOverFiveDecadesAguiarAndHurst2006Dataset, error) {
	var response model.TheAllocationOfTimeOverFiveDecadesAguiarAndHurst2006Dataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) TheWorldsNumberAndShareOfVaccinatedOneYearOlds(ctx context.Context, entity string, year int) (*model.TheWorldsNumberAndShareOfVaccinatedOneYearOldsDataset, error) {
	var response model.TheWorldsNumberAndShareOfVaccinatedOneYearOldsDataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) TimeSpentOnDomesticWorkUn2017AndOecd2014(ctx context.Context, entity string, year int) (*model.TimeSpentOnDomesticWorkUn2017AndOecd2014Dataset, error) {
	var response model.TimeSpentOnDomesticWorkUn2017AndOecd2014Dataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) TimeSpentParticipationTimeAndParticipationRatesEurostat(ctx context.Context, entity string, year int) (*model.TimeSpentParticipationTimeAndParticipationRatesEurostatDataset, error) {
	var response model.TimeSpentParticipationTimeAndParticipationRatesEurostatDataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) TimeThatDoctorsSpendWithAPatientDasHammerAndLeonard2008(ctx context.Context, entity string, year int) (*model.TimeThatDoctorsSpendWithAPatientDasHammerAndLeonard2008Dataset, error) {
	var response model.TimeThatDoctorsSpendWithAPatientDasHammerAndLeonard2008Dataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) TimeUseInFinlandStatisticsFinland(ctx context.Context, entity string, year int) (*model.TimeUseInFinlandStatisticsFinlandDataset, error) {
	var response model.TimeUseInFinlandStatisticsFinlandDataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) TimeUseInSwedenStatisticsSweden(ctx context.Context, entity string, year int) (*model.TimeUseInSwedenStatisticsSwedenDataset, error) {
	var response model.TimeUseInSwedenStatisticsSwedenDataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) Top1percWealthSharesChartbookOfEconomicInequality2017(ctx context.Context, entity string, year int) (*model.Top1percWealthSharesChartbookOfEconomicInequality2017Dataset, error) {
	var response model.Top1percWealthSharesChartbookOfEconomicInequality2017Dataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) TopNetPersonalWealthSharesWid2018(ctx context.Context, entity string, year int) (*model.TopNetPersonalWealthSharesWid2018Dataset, error) {
	var response model.TopNetPersonalWealthSharesWid2018Dataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) TopIncomeSharesWorldWealthAndIncomeDatabase2018(ctx context.Context, entity string, year int) (*model.TopIncomeSharesWorldWealthAndIncomeDatabase2018Dataset, error) {
	var response model.TopIncomeSharesWorldWealthAndIncomeDatabase2018Dataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) TopMarginalIncomeTaxRateReynolds2008(ctx context.Context, entity string, year int) (*model.TopMarginalIncomeTaxRateReynolds2008Dataset, error) {
	var response model.TopMarginalIncomeTaxRateReynolds2008Dataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) TotalEconomyProductivityGrowthOecd(ctx context.Context, entity string, year int) (*model.TotalEconomyProductivityGrowthOecdDataset, error) {
	var response model.TotalEconomyProductivityGrowthOecdDataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) TotalCasesOfPoliomyelitisVirusByCountryAndYearFrom1980OnwardsWho2020(ctx context.Context, entity string, year int) (*model.TotalCasesOfPoliomyelitisVirusByCountryAndYearFrom1980OnwardsWho2020Dataset, error) {
	var response model.TotalCasesOfPoliomyelitisVirusByCountryAndYearFrom1980OnwardsWho2020Dataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) TotalFertilityByRegion2020_2100MediumVariantProjectionUnPopulationDivision2015Revision(ctx context.Context, entity string, year int) (*model.TotalFertilityByRegion20202100MediumVariantProjectionUnPopulationDivision2015RevisionDataset, error) {
	var response model.TotalFertilityByRegion20202100MediumVariantProjectionUnPopulationDivision2015RevisionDataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) TotalGrossOfficialDisbursementsForMedicalResearchAndBasicHeathSectorsOecd(ctx context.Context, entity string, year int) (*model.TotalGrossOfficialDisbursementsForMedicalResearchAndBasicHeathSectorsOecdDataset, error) {
	var response model.TotalGrossOfficialDisbursementsForMedicalResearchAndBasicHeathSectorsOecdDataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) TotalPopulationGapminderUnPopulationDivision(ctx context.Context, entity string, year int) (*model.TotalPopulationGapminderUnPopulationDivisionDataset, error) {
	var response model.TotalPopulationGapminderUnPopulationDivisionDataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) TotalPopulationByBroadAgeGroupBothSexes1950_2100UnPopulationDivision2015(ctx context.Context, entity string, year int) (*model.TotalPopulationByBroadAgeGroupBothSexes19502100UnPopulationDivision2015Dataset, error) {
	var response model.TotalPopulationByBroadAgeGroupBothSexes19502100UnPopulationDivision2015Dataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) TotalValueOfExportsByCountryToWorldPercgdpOwidCalculationsBasedOnFouquinAndHugotCepii2016AndOtherSources(ctx context.Context, entity string, year int) (*model.TotalValueOfExportsByCountryToWorldPercgdpOwidCalculationsBasedOnFouquinAndHugotCepii2016AndOtherSourcesDataset, error) {
	var response model.TotalValueOfExportsByCountryToWorldPercgdpOwidCalculationsBasedOnFouquinAndHugotCepii2016AndOtherSourcesDataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) TourismDataByWorldRegionUnwto2019(ctx context.Context, entity string, year int) (*model.TourismDataByWorldRegionUnwto2019Dataset, error) {
	var response model.TourismDataByWorldRegionUnwto2019Dataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) TradeShareByTypeOfTradeOwidCalculationsBasedOnNberUnitedNationsTradeData1962_2000(ctx context.Context, entity string, year int) (*model.TradeShareByTypeOfTradeOwidCalculationsBasedOnNberUnitedNationsTradeData19622000Dataset, error) {
	var response model.TradeShareByTypeOfTradeOwidCalculationsBasedOnNberUnitedNationsTradeData19622000Dataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) TradeShareWithCapitalAndLaborIntensiveCountriesOwidCalculationsBasedOnFouquinAndHugotCepii2016(ctx context.Context, entity string, year int) (*model.TradeShareWithCapitalAndLaborIntensiveCountriesOwidCalculationsBasedOnFouquinAndHugotCepii2016Dataset, error) {
	var response model.TradeShareWithCapitalAndLaborIntensiveCountriesOwidCalculationsBasedOnFouquinAndHugotCepii2016Dataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) TradeGiovanniAndTenaJunguito2016(ctx context.Context, entity string, year int) (*model.TradeGiovanniAndTenaJunguito2016Dataset, error) {
	var response model.TradeGiovanniAndTenaJunguito2016Dataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) TransistorsPerMicroprocessorRuppAndHorowitz(ctx context.Context, entity string, year int) (*model.TransistorsPerMicroprocessorRuppAndHorowitzDataset, error) {
	var response model.TransistorsPerMicroprocessorRuppAndHorowitzDataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) TreeDensityCrowtherEtAl2015(ctx context.Context, entity string, year int) (*model.TreeDensityCrowtherEtAl2015Dataset, error) {
	var response model.TreeDensityCrowtherEtAl2015Dataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) TropicalDeforestationByCountryOrRegionPendrillEtAl2019(ctx context.Context, entity string, year int) (*model.TropicalDeforestationByCountryOrRegionPendrillEtAl2019Dataset, error) {
	var response model.TropicalDeforestationByCountryOrRegionPendrillEtAl2019Dataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) TrustEurostat(ctx context.Context, entity string, year int) (*model.TrustEurostatDataset, error) {
	var response model.TrustEurostatDataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) TrustWorldValueSurvey(ctx context.Context, entity string, year int) (*model.TrustWorldValueSurveyDataset, error) {
	var response model.TrustWorldValueSurveyDataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) UcdpprioArmedConflictDatasetVersion172DirectFormUcdp(ctx context.Context, entity string, year int) (*model.UcdpprioArmedConflictDatasetVersion172DirectFormUcdpDataset, error) {
	var response model.UcdpprioArmedConflictDatasetVersion172DirectFormUcdpDataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) UkButterflyPopulationsUkOns(ctx context.Context, entity string, year int) (*model.UkButterflyPopulationsUkOnsDataset, error) {
	var response model.UkButterflyPopulationsUkOnsDataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) UkCholeraDeathOverTheLongTermOns(ctx context.Context, entity string, year int) (*model.UkCholeraDeathOverTheLongTermOnsDataset, error) {
	var response model.UkCholeraDeathOverTheLongTermOnsDataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) UkDefenceSpendingUkpublicspendingcom(ctx context.Context, entity string, year int) (*model.UkDefenceSpendingUkpublicspendingcomDataset, error) {
	var response model.UkDefenceSpendingUkpublicspendingcomDataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) UkNominalWageDataPriceDataAndRealWageBankOfEnglandThreeCenturiesOfMacroeconomicData2017(ctx context.Context, entity string, year int) (*model.UkNominalWageDataPriceDataAndRealWageBankOfEnglandThreeCenturiesOfMacroeconomicData2017Dataset, error) {
	var response model.UkNominalWageDataPriceDataAndRealWageBankOfEnglandThreeCenturiesOfMacroeconomicData2017Dataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) UnPopulationDivision2015(ctx context.Context, entity string, year int) (*model.UnPopulationDivision2015Dataset, error) {
	var response model.UnPopulationDivision2015Dataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) UnPopulationDivisionMedianAge2015(ctx context.Context, entity string, year int) (*model.UnPopulationDivisionMedianAge2015Dataset, error) {
	var response model.UnPopulationDivisionMedianAge2015Dataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) UnPopulationDivisionMedianAge2017(ctx context.Context, entity string, year int) (*model.UnPopulationDivisionMedianAge2017Dataset, error) {
	var response model.UnPopulationDivisionMedianAge2017Dataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) UnescoMetadataOnLiteracyUis2017(ctx context.Context, entity string, year int) (*model.UnescoMetadataOnLiteracyUis2017Dataset, error) {
	var response model.UnescoMetadataOnLiteracyUis2017Dataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) UsPublicTrustInGovernmentPewResearchCenter(ctx context.Context, entity string, year int) (*model.UsPublicTrustInGovernmentPewResearchCenterDataset, error) {
	var response model.UsPublicTrustInGovernmentPewResearchCenterDataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) UsCornYieldsUsda2017AndFao2017(ctx context.Context, entity string, year int) (*model.UsCornYieldsUsda2017AndFao2017Dataset, error) {
	var response model.UsCornYieldsUsda2017AndFao2017Dataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) UsFemaleLaborForceParticipation1890_2005Olivetti2013(ctx context.Context, entity string, year int) (*model.UsFemaleLaborForceParticipation18902005Olivetti2013Dataset, error) {
	var response model.UsFemaleLaborForceParticipation18902005Olivetti2013Dataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) UsMaternalMortalityAndFlfpIndexOwid2017(ctx context.Context, entity string, year int) (*model.UsMaternalMortalityAndFlfpIndexOwid2017Dataset, error) {
	var response model.UsMaternalMortalityAndFlfpIndexOwid2017Dataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) UsMeaslesCasesAndDeathsOwid2017(ctx context.Context, entity string, year int) (*model.UsMeaslesCasesAndDeathsOwid2017Dataset, error) {
	var response model.UsMeaslesCasesAndDeathsOwid2017Dataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) UsOpinionOnWivesWorking1936_1998OwidCompilation(ctx context.Context, entity string, year int) (*model.UsOpinionOnWivesWorking19361998OwidCompilationDataset, error) {
	var response model.UsOpinionOnWivesWorking19361998OwidCompilationDataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) UsRevenuePublicSchoolsUsBureauOfTheCensusAndNces2017(ctx context.Context, entity string, year int) (*model.UsRevenuePublicSchoolsUsBureauOfTheCensusAndNces2017Dataset, error) {
	var response model.UsRevenuePublicSchoolsUsBureauOfTheCensusAndNces2017Dataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) UsaConsumerPriceIndexGoodsAndServices1997_2017UsBureauOfLaborStatistics2017(ctx context.Context, entity string, year int) (*model.UsaConsumerPriceIndexGoodsAndServices19972017UsBureauOfLaborStatistics2017Dataset, error) {
	var response model.UsaConsumerPriceIndexGoodsAndServices19972017UsBureauOfLaborStatistics2017Dataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) UsaPolioCasesAndDeaths1910_2010OwidBasedOnUsPublicHealthService1910_1951UsCenterForDiseaseControl1960_2010AndWho2011(ctx context.Context, entity string, year int) (*model.UsaPolioCasesAndDeaths19102010OwidBasedOnUsPublicHealthService19101951UsCenterForDiseaseControl19602010AndWho2011Dataset, error) {
	var response model.UsaPolioCasesAndDeaths19102010OwidBasedOnUsPublicHealthService19101951UsCenterForDiseaseControl19602010AndWho2011Dataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) UsaPatentsGrantedUsPatentAndTrademarkOffice(ctx context.Context, entity string, year int) (*model.UsaPatentsGrantedUsPatentAndTrademarkOfficeDataset, error) {
	var response model.UsaPatentsGrantedUsPatentAndTrademarkOfficeDataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) UkraineRussiaGlobalFoodBasedOnUnFao(ctx context.Context, entity string, year int) (*model.UkraineRussiaGlobalFoodBasedOnUnFaoDataset, error) {
	var response model.UkraineRussiaGlobalFoodBasedOnUnFaoDataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) UkraineRussiaContributionToGlobalFood(ctx context.Context, entity string, year int) (*model.UkraineRussiaContributionToGlobalFoodDataset, error) {
	var response model.UkraineRussiaContributionToGlobalFoodDataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) UnadjustedFemaleMaleHourlyWageRatiosByPercentileBlauKahn2017(ctx context.Context, entity string, year int) (*model.UnadjustedFemaleMaleHourlyWageRatiosByPercentileBlauKahn2017Dataset, error) {
	var response model.UnadjustedFemaleMaleHourlyWageRatiosByPercentileBlauKahn2017Dataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) Under5MortalityRateOurWorldInData(ctx context.Context, entity string, year int) (*model.Under5MortalityRateOurWorldInDataDataset, error) {
	var response model.Under5MortalityRateOurWorldInDataDataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) UnderFiveMortalityRateUnWorldPopulationProspects2015(ctx context.Context, entity string, year int) (*model.UnderFiveMortalityRateUnWorldPopulationProspects2015Dataset, error) {
	var response model.UnderFiveMortalityRateUnWorldPopulationProspects2015Dataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) UnemploymentRateAges25_54ByEducationIlostat2017(ctx context.Context, entity string, year int) (*model.UnemploymentRateAges2554ByEducationIlostat2017Dataset, error) {
	var response model.UnemploymentRateAges2554ByEducationIlostat2017Dataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) UnionDensityQualityOfGovernmentQog2017(ctx context.Context, entity string, year int) (*model.UnionDensityQualityOfGovernmentQog2017Dataset, error) {
	var response model.UnionDensityQualityOfGovernmentQog2017Dataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) UnitedNationsPeacekeeping(ctx context.Context, entity string, year int) (*model.UnitedNationsPeacekeepingDataset, error) {
	var response model.UnitedNationsPeacekeepingDataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) UnitedNationsHumanDevelopmentIndexHdi(ctx context.Context, entity string, year int) (*model.UnitedNationsHumanDevelopmentIndexHdiDataset, error) {
	var response model.UnitedNationsHumanDevelopmentIndexHdiDataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) UrbanAndRuralPopulation1950_2050UnWorldUrbanizationProspects2018(ctx context.Context, entity string, year int) (*model.UrbanAndRuralPopulation19502050UnWorldUrbanizationProspects2018Dataset, error) {
	var response model.UrbanAndRuralPopulation19502050UnWorldUrbanizationProspects2018Dataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) UrbanAndRuralPopulationsInTheUnitedStatesUsCensusBureau2010(ctx context.Context, entity string, year int) (*model.UrbanAndRuralPopulationsInTheUnitedStatesUsCensusBureau2010Dataset, error) {
	var response model.UrbanAndRuralPopulationsInTheUnitedStatesUsCensusBureau2010Dataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) UrbanDefinitionPopulationThresholdUn2018(ctx context.Context, entity string, year int) (*model.UrbanDefinitionPopulationThresholdUn2018Dataset, error) {
	var response model.UrbanDefinitionPopulationThresholdUn2018Dataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) UrbanPopulationLivingInSlumsWbWdi(ctx context.Context, entity string, year int) (*model.UrbanPopulationLivingInSlumsWbWdiDataset, error) {
	var response model.UrbanPopulationLivingInSlumsWbWdiDataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) UrbanizationInTheLongRunOwidBasedOnTheUnWorldUrbanizationProspects2018AndOthers(ctx context.Context, entity string, year int) (*model.UrbanizationInTheLongRunOwidBasedOnTheUnWorldUrbanizationProspects2018AndOthersDataset, error) {
	var response model.UrbanizationInTheLongRunOwidBasedOnTheUnWorldUrbanizationProspects2018AndOthersDataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) UrbanizationShareEuropeanCommissionAtlasOfTheHumanPlanet(ctx context.Context, entity string, year int) (*model.UrbanizationShareEuropeanCommissionAtlasOfTheHumanPlanetDataset, error) {
	var response model.UrbanizationShareEuropeanCommissionAtlasOfTheHumanPlanetDataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) UseOfDifferentSocialMediaSitesByDemographicGroups(ctx context.Context, entity string, year int) (*model.UseOfDifferentSocialMediaSitesByDemographicGroupsDataset, error) {
	var response model.UseOfDifferentSocialMediaSitesByDemographicGroupsDataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) VaccineCoverageAndDiseaseBurdenWho2017(ctx context.Context, entity string, year int) (*model.VaccineCoverageAndDiseaseBurdenWho2017Dataset, error) {
	var response model.VaccineCoverageAndDiseaseBurdenWho2017Dataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) ValueOfGlobalMerchandiseImportsAndExportsFouquinAndHugotCepii2016NationalData(ctx context.Context, entity string, year int) (*model.ValueOfGlobalMerchandiseImportsAndExportsFouquinAndHugotCepii2016NationalDataDataset, error) {
	var response model.ValueOfGlobalMerchandiseImportsAndExportsFouquinAndHugotCepii2016NationalDataDataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) ViolentVictimizationUsBureauOfJusticeStatistics2017(ctx context.Context, entity string, year int) (*model.ViolentVictimizationUsBureauOfJusticeStatistics2017Dataset, error) {
	var response model.ViolentVictimizationUsBureauOfJusticeStatistics2017Dataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) ViolentDeathsInConflictsAndOneSidedViolenceSince1989ByRegionAndTypeOfViolenceUcdp2022(ctx context.Context, entity string, year int) (*model.ViolentDeathsInConflictsAndOneSidedViolenceSince1989ByRegionAndTypeOfViolenceUcdp2022Dataset, error) {
	var response model.ViolentDeathsInConflictsAndOneSidedViolenceSince1989ByRegionAndTypeOfViolenceUcdp2022Dataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) ViolentDisciplineInTheUsUsGeneralSocialSurvey2017(ctx context.Context, entity string, year int) (*model.ViolentDisciplineInTheUsUsGeneralSocialSurvey2017Dataset, error) {
	var response model.ViolentDisciplineInTheUsUsGeneralSocialSurvey2017Dataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) ViolentDisciplineUnicef2017(ctx context.Context, entity string, year int) (*model.ViolentDisciplineUnicef2017Dataset, error) {
	var response model.ViolentDisciplineUnicef2017Dataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) VolcanicEruptionDeathsNgdcnoaa(ctx context.Context, entity string, year int) (*model.VolcanicEruptionDeathsNgdcnoaaDataset, error) {
	var response model.VolcanicEruptionDeathsNgdcnoaaDataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) WagesInTheManufacturingSectorVsSeveralFoodPricesInTheUsUsBureauOfLaborStatistics2013(ctx context.Context, entity string, year int) (*model.WagesInTheManufacturingSectorVsSeveralFoodPricesInTheUsUsBureauOfLaborStatistics2013Dataset, error) {
	var response model.WagesInTheManufacturingSectorVsSeveralFoodPricesInTheUsUsBureauOfLaborStatistics2013Dataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) WaterAndSanitationWhoWash2021(ctx context.Context, entity string, year int) (*model.WaterAndSanitationWhoWash2021Dataset, error) {
	var response model.WaterAndSanitationWhoWash2021Dataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) WaterResourcesByContinentFaoAquastat(ctx context.Context, entity string, year int) (*model.WaterResourcesByContinentFaoAquastatDataset, error) {
	var response model.WaterResourcesByContinentFaoAquastatDataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) WaterWithdrawalsAndConsumptionAquastat(ctx context.Context, entity string, year int) (*model.WaterWithdrawalsAndConsumptionAquastatDataset, error) {
	var response model.WaterWithdrawalsAndConsumptionAquastatDataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) WattsPerMipsKurzweil(ctx context.Context, entity string, year int) (*model.WattsPerMipsKurzweilDataset, error) {
	var response model.WattsPerMipsKurzweilDataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) WealthTotalByComponentForVariousCountryGroupingsWorldBank2017(ctx context.Context, entity string, year int) (*model.WealthTotalByComponentForVariousCountryGroupingsWorldBank2017Dataset, error) {
	var response model.WealthTotalByComponentForVariousCountryGroupingsWorldBank2017Dataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) WealthAsPercentNationalIncomeByWealthTypePiketty2014(ctx context.Context, entity string, year int) (*model.WealthAsPercentNationalIncomeByWealthTypePiketty2014Dataset, error) {
	var response model.WealthAsPercentNationalIncomeByWealthTypePiketty2014Dataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) WealthPerCapitaByComponentByCountryWorldBank2017(ctx context.Context, entity string, year int) (*model.WealthPerCapitaByComponentByCountryWorldBank2017Dataset, error) {
	var response model.WealthPerCapitaByComponentByCountryWorldBank2017Dataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) WealthPerCapitaByComponentForVariousCountryGroupingsWorldBank2017(ctx context.Context, entity string, year int) (*model.WealthPerCapitaByComponentForVariousCountryGroupingsWorldBank2017Dataset, error) {
	var response model.WealthPerCapitaByComponentForVariousCountryGroupingsWorldBank2017Dataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) WeatherFatalityRatesInTheUsOwidBasedOnNoaaAndLopezHolleAndPopulationData(ctx context.Context, entity string, year int) (*model.WeatherFatalityRatesInTheUsOwidBasedOnNoaaAndLopezHolleAndPopulationDataDataset, error) {
	var response model.WeatherFatalityRatesInTheUsOwidBasedOnNoaaAndLopezHolleAndPopulationDataDataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) WeeklyHomeProductionHoursInTheUsaRamey2009AndRameyFrancis2009(ctx context.Context, entity string, year int) (*model.WeeklyHomeProductionHoursInTheUsaRamey2009AndRameyFrancis2009Dataset, error) {
	var response model.WeeklyHomeProductionHoursInTheUsaRamey2009AndRameyFrancis2009Dataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) WellcomeGlobalMonitorTrust(ctx context.Context, entity string, year int) (*model.WellcomeGlobalMonitorTrustDataset, error) {
	var response model.WellcomeGlobalMonitorTrustDataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) WhaleCatchRochaEtAlIwc(ctx context.Context, entity string, year int) (*model.WhaleCatchRochaEtAlIwcDataset, error) {
	var response model.WhaleCatchRochaEtAlIwcDataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) WhaleCatchByDecadeRochaEtAlAndIwc(ctx context.Context, entity string, year int) (*model.WhaleCatchByDecadeRochaEtAlAndIwcDataset, error) {
	var response model.WhaleCatchByDecadeRochaEtAlAndIwcDataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) WhalePopulationsPershingEtAl2010(ctx context.Context, entity string, year int) (*model.WhalePopulationsPershingEtAl2010Dataset, error) {
	var response model.WhalePopulationsPershingEtAl2010Dataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) WheatPricesLongRunInEnglandMakridakisEtAl1997(ctx context.Context, entity string, year int) (*model.WheatPricesLongRunInEnglandMakridakisEtAl1997Dataset, error) {
	var response model.WheatPricesLongRunInEnglandMakridakisEtAl1997Dataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) WhoAmericansSpendTimeWithAmericanTimeUseSurvey2009_2019(ctx context.Context, entity string, year int) (*model.WhoAmericansSpendTimeWithAmericanTimeUseSurvey20092019Dataset, error) {
	var response model.WhoAmericansSpendTimeWithAmericanTimeUseSurvey20092019Dataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) WildfireDataInTheUsNifc(ctx context.Context, entity string, year int) (*model.WildfireDataInTheUsNifcDataset, error) {
	var response model.WildfireDataInTheUsNifcDataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) WomensEconomicOpportunity2012EconomistIntelligenceUnit2012(ctx context.Context, entity string, year int) (*model.WomensEconomicOpportunity2012EconomistIntelligenceUnit2012Dataset, error) {
	var response model.WomensEconomicOpportunity2012EconomistIntelligenceUnit2012Dataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) WomensWeeklyEarningsAsAPercentageOfMensBureauOfLaborStatistics2017(ctx context.Context, entity string, year int) (*model.WomensWeeklyEarningsAsAPercentageOfMensBureauOfLaborStatistics2017Dataset, error) {
	var response model.WomensWeeklyEarningsAsAPercentageOfMensBureauOfLaborStatistics2017Dataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) WomensPoliticalRepresentationUsingPaxtonEtAl2006Ipu2017AndWdi2017(ctx context.Context, entity string, year int) (*model.WomensPoliticalRepresentationUsingPaxtonEtAl2006Ipu2017AndWdi2017Dataset, error) {
	var response model.WomensPoliticalRepresentationUsingPaxtonEtAl2006Ipu2017AndWdi2017Dataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) WorkingHoursDataHubermanAndMinns2007(ctx context.Context, entity string, year int) (*model.WorkingHoursDataHubermanAndMinns2007Dataset, error) {
	var response model.WorkingHoursDataHubermanAndMinns2007Dataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) WorldBankIncomeThresholdsWorldBank2017(ctx context.Context, entity string, year int) (*model.WorldBankIncomeThresholdsWorldBank2017Dataset, error) {
	var response model.WorldBankIncomeThresholdsWorldBank2017Dataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) WorldBankEducationDatasetWorldBank2015(ctx context.Context, entity string, year int) (*model.WorldBankEducationDatasetWorldBank2015Dataset, error) {
	var response model.WorldBankEducationDatasetWorldBank2015Dataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) WorldGdpIn2011IntMoneyOwidBasedOnWorldBankMaddison2017(ctx context.Context, entity string, year int) (*model.WorldGdpIn2011IntMoneyOwidBasedOnWorldBankMaddison2017Dataset, error) {
	var response model.WorldGdpIn2011IntMoneyOwidBasedOnWorldBankMaddison2017Dataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) WorldHappinessReport2022(ctx context.Context, entity string, year int) (*model.WorldHappinessReport2022Dataset, error) {
	var response model.WorldHappinessReport2022Dataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) WorldPopulationByPoliticalRegimeTheyLiveInOwid2016(ctx context.Context, entity string, year int) (*model.WorldPopulationByPoliticalRegimeTheyLiveInOwid2016Dataset, error) {
	var response model.WorldPopulationByPoliticalRegimeTheyLiveInOwid2016Dataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) WorldPovertyClock(ctx context.Context, entity string, year int) (*model.WorldPovertyClockDataset, error) {
	var response model.WorldPovertyClockDataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) WorldPressFreedomIndexReportersSansFrontieres2022(ctx context.Context, entity string, year int) (*model.WorldPressFreedomIndexReportersSansFrontieres2022Dataset, error) {
	var response model.WorldPressFreedomIndexReportersSansFrontieres2022Dataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) WorldRegionsAccordingToTheWorldBank(ctx context.Context, entity string, year int) (*model.WorldRegionsAccordingToTheWorldBankDataset, error) {
	var response model.WorldRegionsAccordingToTheWorldBankDataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) WorldConflictDeathRateSince1989VariousSources(ctx context.Context, entity string, year int) (*model.WorldConflictDeathRateSince1989VariousSourcesDataset, error) {
	var response model.WorldConflictDeathRateSince1989VariousSourcesDataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) WorldConflictDeathsVariousSources(ctx context.Context, entity string, year int) (*model.WorldConflictDeathsVariousSourcesDataset, error) {
	var response model.WorldConflictDeathsVariousSourcesDataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) YearOfMaternalAndNeonatalTetanusMntEliminationWhoAndKiwanis2018(ctx context.Context, entity string, year int) (*model.YearOfMaternalAndNeonatalTetanusMntEliminationWhoAndKiwanis2018Dataset, error) {
	var response model.YearOfMaternalAndNeonatalTetanusMntEliminationWhoAndKiwanis2018Dataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) YearOfLastRinderpestCaseOie2018(ctx context.Context, entity string, year int) (*model.YearOfLastRinderpestCaseOie2018Dataset, error) {
	var response model.YearOfLastRinderpestCaseOie2018Dataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) YearOfLastRecordedWildPoliomyelitisVirusWhoGpei2017(ctx context.Context, entity string, year int) (*model.YearOfLastRecordedWildPoliomyelitisVirusWhoGpei2017Dataset, error) {
	var response model.YearOfLastRecordedWildPoliomyelitisVirusWhoGpei2017Dataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) YearOfSmallpoxEradicationByCountryWho1988(ctx context.Context, entity string, year int) (*model.YearOfSmallpoxEradicationByCountryWho1988Dataset, error) {
	var response model.YearOfSmallpoxEradicationByCountryWho1988Dataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) YearsOfSchoolingBasedOnLeeLee2016BarroLee2018AndUndp2018(ctx context.Context, entity string, year int) (*model.YearsOfSchoolingBasedOnLeeLee2016BarroLee2018AndUndp2018Dataset, error) {
	var response model.YearsOfSchoolingBasedOnLeeLee2016BarroLee2018AndUndp2018Dataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) YougovImperialCovid19BehaviorTracker(ctx context.Context, entity string, year int) (*model.YougovImperialCovid19BehaviorTrackerDataset, error) {
	var response model.YougovImperialCovid19BehaviorTrackerDataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

func (r *queryResolver) YouthMortalityRateUnIgme2018(ctx context.Context, entity string, year int) (*model.YouthMortalityRateUnIgme2018Dataset, error) {
	var response model.YouthMortalityRateUnIgme2018Dataset
	_, err := r.Fetcher(ctx, entity, year, &response)
	return &response, err
}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
