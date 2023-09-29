#!/bin/bash

set -xe

echo "Today is " `date`

echo -e "\nEnter a Region/LOCATION name"
read LOCATION

echo -e "\nEnter a cluster name"
read CLUSTER

echo -e "\nEnter the ResourceGroup"
read RESOURCEGROUP

echo -e "\nEnter a Cluster version"
read VERSION

echo -e "\n Your cluster will be created in $LOCATION region, ResourceGroup $RESOURCEGROUP, with version $VERSION and the Cluster name will be: $CLUSTER"


echo "Creating Resource Group"
az group create \
--name $RESOURCEGROUP \
--location $LOCATION

if [ $? -ne 0 ]; then
    echo "Error occurred."
fi

echo -e "\n==================================================================="

echo "Creating VNET"
az network vnet create \
--resource-group $RESOURCEGROUP \
--name aro-vnet \
--address-prefixes 10.0.0.0/22

if [ $? -ne 0 ]; then
    echo "Error occurred."
fi

echo -e "\n ==================================================================="
echo "Creating Master Subnet"
az network vnet subnet create \
--resource-group $RESOURCEGROUP \
--vnet-name aro-vnet \
--name master-subnet \
--address-prefixes 10.0.0.0/23

if [ $? -ne 0 ]; then
    echo "Error occurred."
fi

echo -e "\n ==================================================================="
echo -e "\nCreating Worker Subnet"
az network vnet subnet create \
--resource-group $RESOURCEGROUP \
--vnet-name aro-vnet \
--name worker-subnet \
--address-prefixes 10.0.2.0/23

if [ $? -ne 0 ]; then
    echo "Error occurred."
fi

echo -e "\n==================================================================="
echo -e "\nCreating a Public API visibility cluster with version $VERSION"
az aro create \
--resource-group $RESOURCEGROUP \
--name $CLUSTER \
--vnet aro-vnet \
--master-subnet master-subnet \
--worker-subnet worker-subnet \
--version $VERSION

if [ $? -ne 0 ]; then
    echo "Error occurred."
fi

