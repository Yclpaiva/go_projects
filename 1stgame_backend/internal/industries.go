package internal

type Industry struct {
	Name        string
	Production  map[string]float64
	Consumption map[string]float64
	Price       int
	CardTitle   string
	CardImage   string
}

/*
	{ 'id': 'woodIndustry', 'price': 5_000, 'cardTitle': 'Wood Industry', 'cardImage': `${industriesImageSource}woodIndustry.jpg` },

{ 'id': 'ironMine', 'price': 10_000, 'cardTitle': 'Iron Mine', 'cardImage': `${industriesImageSource}ironMine.jpg` },
{ 'id': 'farmlands', 'price': 3_000, 'cardTitle': 'Farmlands', 'cardImage': `${industriesImageSource}farmlands.jpg` },
{ 'id': 'foodFactory', 'price': 50_000, 'cardTitle': 'Food Factory', 'cardImage': `${industriesImageSource}foodFactory.jpg` },
{ 'id': 'livestockFarm', 'price': 20_000, 'cardTitle': 'Livestock Farm', 'cardImage': `${industriesImageSource}livestockFarm.jpg` },
{ 'id': 'fishingPort', 'price': 35_000, 'cardTitle': 'Fishing Port', 'cardImage': `${industriesImageSource}fishingPort.jpg` },
{ 'id': 'mechanicalPartsIndustry', 'price': 100_000, 'cardTitle': 'Mechanical Parts Industry', 'cardImage': `${industriesImageSource}mechanicalPartsIndustry.jpg` },
{ 'id': 'coalMine', 'price': 10_000, 'cardTitle': 'Coal Mine', 'cardImage': `${industriesImageSource}coalMine.jpeg` },
{ 'id': 'coalPowerplant', 'price': 100_000, 'cardTitle': 'Coal Powerplay', 'cardImage': `${industriesImageSource}coalPowerplant.jpg` },
{ 'id': 'copperMine', 'price': 30_000, 'cardTitle': 'Copper Mine', 'cardImage': `${industriesImageSource}copperMine.jpg` },
{ 'id': 'cottonFarm', 'price': 5_000, 'cardTitle': 'Cotton Farm', 'cardImage': `${industriesImageSource}cottonFarm.jpg` },
{ 'id': 'paperFactory', 'price': 100_000, 'cardTitle': 'Paper Factory', 'cardImage': `${industriesImageSource}paperFactory.jpg` },
{ 'id': 'goldMine', 'price': 1_000_000, 'cardTitle': 'Gold Mine', 'cardImage': `${industriesImageSource}goldMine.jpg` },
{ 'id': 'university', 'price': 500_000, 'cardTitle': 'University', 'cardImage': `${industriesImageSource}university.jpg` },
*/
var industries = []Industry{
	{
		Name: "woodIndustry",
		Production: map[string]float64{
			"wood": 1,
		},
		Price:     5_000,
		CardTitle: "Wood Industry",
		CardImage: `woodIndustry.jpg`,
	},
	{
		Name: "farmlands",
		Production: map[string]float64{
			"grain": 1,
		},
		Price:     3_000,
		CardTitle: "Farmlands",
		CardImage: `farmlands.jpg`,
	},
	{
		Name: "ironMine",
		Production: map[string]float64{
			"iron": 1,
		},
		Price:     10_000,
		CardTitle: "Iron Mine",
		CardImage: `ironMine.jpg`,
	},
	{
		Name: "foodFactory",
		Production: map[string]float64{
			"processedFood": 1,
		},
		Consumption: map[string]float64{
			"grain": 5,
		},
		Price:     50_000,
		CardTitle: "Food Factory",
		CardImage: `foodFactory.jpg`,
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
		Price:     20_000,
		CardTitle: "Livestock Farm",
		CardImage: `livestockFarm.jpg`,
	},
	{
		Name: "fishingPort",
		Production: map[string]float64{
			"protein": 1,
		},
		Price:     35_000,
		CardTitle: "Fishing Port",
		CardImage: `fishingPort.jpg`,
	},
	{
		Name: "mechanicalPartsIndustry",
		Production: map[string]float64{
			"mechanicalParts": 1,
		},
		Consumption: map[string]float64{
			"iron": 5,
		},
		Price:     100_000,
		CardTitle: "Mechanical Parts Industry",
		CardImage: `mechanicalPartsIndustry.jpg`,
	},
	{
		Name: "coalMine",
		Production: map[string]float64{
			"coal": 1,
		},
		Price:     10_000,
		CardTitle: "Coal Mine",
		CardImage: `coalMine.jpeg`,
	},
	{
		Name: "coalPowerplant",
		Production: map[string]float64{
			"electricity": 15,
		},
		Consumption: map[string]float64{
			"coal": 20,
		},
		Price:     100_000,
		CardTitle: "Coal Powerplay",
		CardImage: `coalPowerplant.jpg`,
	},
	{
		Name: "goldMine",
		Production: map[string]float64{
			"gold": 0.1,
		},
		Price:     1_000_000,
		CardTitle: "Gold Mine",
		CardImage: `goldMine.jpg`,
	},
	{
		Name: "copperMine",
		Production: map[string]float64{
			"copper": 1,
		},
		Price:     30_000,
		CardTitle: "Copper Mine",
		CardImage: `copperMine.jpg`,
	},
	{
		Name: "cottonFarm",
		Production: map[string]float64{
			"cotton": 1,
		},
		Price:     5_000,
		CardTitle: "Cotton Farm",
		CardImage: `cottonFarm.jpg`,
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
		Price:     100_000,
		CardTitle: "Paper Factory",
		CardImage: `paperFactory.jpg`,
	},
	{
		Name: "university",
		Production: map[string]float64{
			"research": 2,
		},
		Consumption: map[string]float64{
			"paper": 5,
		},
		Price:     500_000,
		CardTitle: "University",
		CardImage: `university.jpg`,
	},
	{
		Name: "siliconIndustry",
		Production: map[string]float64{
			"silicon": 1,
		},
		Price:     500_000,
		CardTitle: "Silicon Industry",
		CardImage: `siliconIndustry.jpg`,
	},
	{
		Name: "sulfurMine",
		Production: map[string]float64{
			"sulfur": 1,
		},
		Price:     100_000,
		CardTitle: "Sulfur Mine",
		CardImage: `sulfurMine.jpg`,
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
		Price:     500_000,
		CardTitle: "Chemical Industry",
		CardImage: `chemicalIndustry.jpg`,
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
		Price:     1_000_000,
		CardTitle: "Components Industry",
		CardImage: `componentsIndustry.jpg`,
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
		Price:     2_000_000,
		CardTitle: "Basic Circuit Factory",
		CardImage: `basicCircuitFactory.jpg`,
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
		Price:     5_000_000,
		CardTitle: "Basic Chips Factory",
		CardImage: `basicChipsFactory.jpg`,
	},
}
