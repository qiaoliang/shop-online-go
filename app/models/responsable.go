package models

type Responsable struct {
    Code        uint    `json:"code"`
    Result    []Banner  `json:"data"`
    Msg     string      `json:"msg"`
}