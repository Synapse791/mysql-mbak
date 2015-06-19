package main

import ()

type Config struct {
	Connections []ConnectionConfig
	S3Config	S3Config
}

type ConnectionConfig struct {
    S3Bucket    string      `json:"s3_upload"`
    Hostname    string      `json:"hostname"`
    Username    string      `json:"username"`
    Password    string      `json:"password"`
    Databases   []string    `json:"databases"`
}

type S3Config struct {
    Region          string  `json:"region"`
    AccessKey       string  `json:"access_key"`
    ClientSecret    string  `json:"client_secret"`
}