module challenge.kamontat.net/shortest-path

go 1.15

replace challenge.kamontat.net/crea-logic => ../logic

replace challenge.kamontat.net/crea-datasource => ../datasource

replace challenge.kamontat.net/crea-model => ../model

replace challenge.kamontat.net/crea-formatter => ../formatter

require (
	challenge.kamontat.net/crea-datasource v0.0.0-00010101000000-000000000000
	challenge.kamontat.net/crea-formatter v0.0.0-00010101000000-000000000000
	challenge.kamontat.net/crea-logic v0.0.0-00010101000000-000000000000
	challenge.kamontat.net/crea-model v0.0.0-00010101000000-000000000000 // indirect
)
