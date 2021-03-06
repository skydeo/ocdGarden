package main

import (
	"opencoredata.org/ocdGarden/CSDCOGraphBuilder/csvproc"
	"opencoredata.org/ocdGarden/CSDCOGraphBuilder/igsnGraph"
	"opencoredata.org/ocdGarden/CSDCOGraphBuilder/pkggrapher"
)

// "strconv"

// "opencoredata.org/ocdCSDCO/metaimport/connectors"

// Yes, in that file the “Hole ID” is a unique identifier for the borehole.
// Cores from each of those holes will have as their unique identifier the Hole ID
//  plus a suffix for the core. And sections of cores will have as their unique
//  identifier the Core ID plus a suffix for the section.
func main() {
	csvproc.BuildGraph()
	pkggrapher.PKGGrapher()
	igsnGraph.BuildGraph()
	// pkgindexer.PKGIndex()
}
