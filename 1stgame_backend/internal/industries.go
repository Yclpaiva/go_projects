package internal

type Industry struct {
	Name        string
	Production  map[string]float64
	Consumption map[string]float64
}

var industries = []Industry{
	{
		Name: "woodIndustry",
		Production: map[string]float64{
			"wood": 1,
		},
	},
	{
		Name: "farmlands",
		Production: map[string]float64{
			"grain": 1,
		},
	},
	{
		Name: "ironMine",
		Production: map[string]float64{
			"iron": 1,
		},
	},
	{
		Name: "foodFactory",
		Production: map[string]float64{
			"processedFood": 1,
		},
		Consumption: map[string]float64{
			"grain": 5,
		},
	},
	{
		Name: "livestockFarm",
		Production: map[string]float64{
			"protein": 1,
		},
		Consumption: map[string]float64{
			"grain": 4,
			"wood":  1,
		},
	},
	{
		Name: "fishingPort",
		Production: map[string]float64{
			"protein": 1,
		},
	},
	{
		Name: "mechanicalPartsIndustry",
		Production: map[string]float64{
			"mechanicalParts": 1,
		},
		Consumption: map[string]float64{
			"iron": 5,
		},
	},
	{
		Name: "coalMine",
		Production: map[string]float64{
			"coal": 1,
		},
	},
	{
		Name: "coalPowerplant",
		Production: map[string]float64{
			"electricity": 15,
		},
		Consumption: map[string]float64{
			"coal": 20,
		},
	},
	{
		Name: "goldMine",
		Production: map[string]float64{
			"gold": 0.1,
		},
	},
	{
		Name: "copperMine",
		Production: map[string]float64{
			"copper": 1,
		},
	},
	{
		Name: "cottonFarm",
		Production: map[string]float64{
			"cotton": 1,
		},
	},
	{
		Name: "paperFactory",
		Production: map[string]float64{
			"paper": 2,
		},
		Consumption: map[string]float64{
			"cotton": 5,
			"wood":   5,
		},
	},
	{
		Name: "university",
		Production: map[string]float64{
			"research": 2,
		},
		Consumption: map[string]float64{
			"paper": 5,
		},
	},
	{
		Name: "siliconIndustry",
		Production: map[string]float64{
			"silicon": 1,
		},
	},
	{
		Name: "sulfurMine",
		Production: map[string]float64{
			"sulfur": 1,
		},
	},
	{
		Name: "chemicalIndustry",
		Production: map[string]float64{
			"chemicals": 0.2,
		},
		Consumption: map[string]float64{
			"sulfur":      1,
			"electricity": 2,
			"grain":       1,
		},
	},
	{
		Name: "componentsIndustry",
		Production: map[string]float64{
			"components": 1,
		},
		Consumption: map[string]float64{
			"chemicals": 1,
			"copper":    10,
		},
	},
	{
		Name: "BasicCircuitFactory",
		Production: map[string]float64{
			"basicCircuit": 0.2,
		},
		Consumption: map[string]float64{
			"silicon":     1,
			"copper":      1,
			"iron":        1,
			"electricity": 1,
			"components":  1,
		},
	},
	{
		Name: "BasicChipsFactory",
		Production: map[string]float64{
			"basicChip": 0.1,
		},
		Consumption: map[string]float64{
			"BasicCircuit": 0.2,
			"components":   1,
			"silicon":      3,
			"electricity":  1,
			"chemicals":    1,
		},
	},
}
