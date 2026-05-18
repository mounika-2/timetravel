package service

import (
	"fmt"
)

type GeminiService struct {
	APIKey string
}

func NewGeminiService(apiKey string) *GeminiService {
	return &GeminiService{
		APIKey: apiKey,
	}
}

func (g *GeminiService) AnalyzeChanges(
	fromData map[string]string,
	toData map[string]string,
) (string, error) {

	changes := []string{}

	// detect changed + removed fields
	for key, oldValue := range fromData {

		newValue, exists := toData[key]

		if !exists {

			changes = append(
				changes,
				fmt.Sprintf(
					"%s was removed",
					key,
				),
			)

			continue
		}

		if oldValue != newValue {

			changes = append(
				changes,
				fmt.Sprintf(
					"%s changed from '%s' to '%s'",
					key,
					oldValue,
					newValue,
				),
			)
		}
	}

	// detect newly added fields
	for key, value := range toData {

		if _, exists := fromData[key]; !exists {

			changes = append(
				changes,
				fmt.Sprintf(
					"%s added with value '%s'",
					key,
					value,
				),
			)
		}
	}

	// no changes
	if len(changes) == 0 {
		return "No material underwriting changes detected.", nil
	}

	analysis := "UNDERWRITING CHANGE ANALYSIS\n\n"

	for _, c := range changes {
		analysis += "• " + c + "\n"
	}

	analysis += "\nRisk Impact Assessment:\n"

	// simple underwriting intelligence

	if employees, ok := toData["employee_count"]; ok {

		analysis += fmt.Sprintf(
			"• Current employee count reported as %s.\n",
			employees,
		)
	}

	if payroll, ok := toData["annual_payroll"]; ok {

		analysis += fmt.Sprintf(
			"• Payroll exposure currently reported as %s.\n",
			payroll,
		)
	}

	if limit, ok := toData["general_liability_limit"]; ok {

		analysis += fmt.Sprintf(
			"• General liability limit currently reported as %s.\n",
			limit,
		)
	}

	if _, ok := toData["hazardous_materials"]; ok {

		analysis +=
			"• Hazardous materials exposure may increase underwriting risk.\n"
	}

	if _, ok := toData["delivery_services"]; ok {

		analysis +=
			"• Delivery operations may introduce additional auto liability exposure.\n"
	}

	if _, ok := toData["overnight_baking"]; ok {

		analysis +=
			"• Overnight operations may impact fire and workers compensation exposure.\n"
	}

	analysis +=
		"\nRecommended underwriting review advised."

	return analysis, nil
}
