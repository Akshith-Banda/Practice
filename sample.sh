#!/bin/bash

echo 'Hello, World!'


export appId="d5dd7f63-6f76-481b-adc2-d82b1606cef8"
export clientsecret="rLR8Q~byY~OyPORZONEm11suTTSk9DmbvO7FMa._"
export tenant="97f998a5-f0f3-4588-8ba3-f956f2f2dcd6"

export region="eastus2"
export resourcegroup="akshith-test"
export uniqueaksclustername="cape-cluster-$resourcegroup"

az login --service-principal -u $appid -p $clientsecret --tenant $tenant


#Create Azrue resource group for our test setup.

az group create --name $resourcegroup --location $region

#Creating the actual AKS Cluster 

az aks create --resource-group $resourcegroup \
    --name $uniqueaksclustername \
    --location $region \
    --kubernetes-version 1.22.6 \
    --generate-ssh-keys \
    --node-count 2 \
    --node-vm-size Standard_D2_v2

#Get cluster access to the local cloud shell
az aks get-credentials --resource-group $resourcegroup  --name $uniqueaksclustername

echo 'Bye, World!'
