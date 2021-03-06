package venom

import (
	"encoding/xml"
)

// Tests contains all informations about tests in a pipeline build
type Tests struct {
	XMLName      xml.Name    `xml:"testsuites" json:"-" yaml:"-"`
	Total        int         `xml:"-" json:"total"`
	TotalOK      int         `xml:"-" json:"ok"`
	TotalKO      int         `xml:"-" json:"ko"`
	TotalSkipped int         `xml:"-" json:"skipped"`
	TestSuites   []TestSuite `xml:"testsuite" json:"test_suites"`
}

// TestSuite is a single JUnit test suite which may contain many
// testcases.
type TestSuite struct {
	XMLName    xml.Name   `xml:"testsuite" json:"-" yaml:"-"`
	Disabled   int        `xml:"disabled,attr,omitempty" json:"disabled" yaml:"-"`
	Errors     int        `xml:"errors,attr,omitempty" json:"errors" yaml:"-"`
	Failures   int        `xml:"failures,attr,omitempty" json:"failures" yaml:"-"`
	Hostname   string     `xml:"hostname,attr,omitempty" json:"hostname" yaml:"-"`
	ID         string     `xml:"id,attr,omitempty" json:"id" yaml:"-"`
	Name       string     `xml:"name,attr" json:"name" yaml:"name"`
	Package    string     `xml:"package,attr,omitempty" json:"package" yaml:"-"`
	Properties []Property `xml:"-" json:"properties" yaml:"-"`
	Skipped    int        `xml:"skipped,attr,omitempty" json:"skipped" yaml:"skipped,omitempty"`
	Total      int        `xml:"tests,attr" json:"total" yaml:"total,omitempty"`
	TestCases  []TestCase `xml:"testcase" json:"tests" yaml:"testcases"`
	Time       string     `xml:"time,attr,omitempty" json:"time" yaml:"-"`
	Timestamp  string     `xml:"timestamp,attr,omitempty" json:"timestamp" yaml:"-"`
}

// Property represents a key/value pair used to define properties.
type Property struct {
	XMLName xml.Name `xml:"property" json:"-" yaml:"-"`
	Name    string   `xml:"name,attr" json:"name" yaml:"-"`
	Value   string   `xml:"value,attr" json:"value" yaml:"-"`
}

// TestCase is a single test case with its result.
type TestCase struct {
	XMLName    xml.Name    `xml:"testcase" json:"-" yaml:"-"`
	Assertions string      `xml:"assertions,attr,omitempty" json:"assertions" yaml:"-"`
	Classname  string      `xml:"classname,attr,omitempty" json:"classname" yaml:"-"`
	Errors     []Failure   `xml:"error,omitempty" json:"errors" yaml:"errors,omitempty"`
	Failures   []Failure   `xml:"failure,omitempty" json:"failures" yaml:"failures,omitempty"`
	Name       string      `xml:"name,attr" json:"name" yaml:"name"`
	Skipped    int         `xml:"skipped,attr,omitempty" json:"skipped" yaml:"skipped,omitempty"`
	Status     string      `xml:"status,attr,omitempty" json:"status" yaml:"status,omitempty"`
	Systemout  InnerResult `xml:"system-out,omitempty" json:"systemout" yaml:"systemout,omitempty"`
	Systemerr  InnerResult `xml:"system-err,omitempty" json:"systemerr" yaml:"systemerr,omitempty"`
	Time       string      `xml:"time,attr,omitempty" json:"time" yaml:"time,omitempty"`
	TestSteps  []TestStep  `xml:"-" json:"steps" yaml:"steps"`
}

// TestStep represents a testStep
type TestStep map[string]interface{}

// Failure contains data related to a failed test.
type Failure struct {
	Value   string `xml:",innerxml" json:"value" yaml:"value,omitempty"`
	Type    string `xml:"type,attr,omitempty" json:"type" yaml:"type,omitempty"`
	Message string `xml:"message,attr,omitempty" json:"message" yaml:"message,omitempty"`
}

// InnerResult is used by TestCase
type InnerResult struct {
	Value string `xml:",innerxml" json:"value" yaml:"value"`
}
