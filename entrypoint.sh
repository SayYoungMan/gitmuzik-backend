#!/bin/bash
export API_KEY=$(cat apikey.txt)
mv credentials ~/.aws/credentials

./main