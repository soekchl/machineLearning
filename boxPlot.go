// 箱型图
package main

import (
	"log"
	"os"

	"github.com/kniren/gota/dataframe"
	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
)

func boxPlot(filePath string) {
	//Open the csv file.
	irisFile, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}
	defer irisFile.Close()

	// Create a dataframe from the CSV file.
	irisDF := dataframe.ReadCSV(irisFile)

	log.Println(irisDF)

	// Create the plot and set its title and axis label.
	p, err := plot.New()
	if err != nil {
		log.Fatal(err)
	}

	p.Title.Text = "Box plots"
	p.Y.Label.Text = "Values"

	// Create the box for our data.
	w := vg.Points(50)

	// Create a box plot for each of the feature columns in the dataset.
	for idx, colName := range irisDF.Names() {
		// If the column is one of the feature columns, let`s create
		// a histogram of the values.
		if colName == "Iris-setosa" {
			continue
		}

		//Create a plotter.Values value and fill it with the
		//values from the respective column of the dataframe.
		v := make(plotter.Values, irisDF.Nrow())

		for i, floatVal := range irisDF.Col(colName).Float() {
			v[i] = floatVal
		}

		// Add the data to the plot.
		b, err := plotter.NewBoxPlot(w, float64(idx), v)
		if err != nil {
			log.Fatal(err)
		}
		p.Add(b)
	}

	// Set the X axis of the plot to nominal with
	// the given names for x=0, x=1, etc.
	p.NominalX("sepal_length", "sepal_width", "petal_length", "petal_width")
	if err = p.Save(6*vg.Inch, 8*vg.Inch, outDir+"boxplots.png"); err != nil {
		log.Fatal(err)
	}
}
