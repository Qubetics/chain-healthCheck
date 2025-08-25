#!/bin/bash

sudo su -c  "echo '[Unit]
Description=Chain HealthCheck Go Service
Wants=network-online.target
After=network-online.target
[Service]
User=$(whoami)
Group=$(whoami)
Type=simple
ExecStart=$(pwd)/chain-healthcheck
Restart=always
RestartSec=3
LimitNOFILE=4096
EnvironmentFile=$(pwd)/.env
[Install]
WantedBy=multi-user.target'> /etc/systemd/system/chainHealthCheck.service"

sudo systemctl daemon-reload
sudo systemctl enable chainHealthCheck.service
sudo systemctl start chainHealthCheck.service