#!/bin/sh
npm run build
echo "Linking local copy"
npm link

# property_manager
echo "Linking property manager to use local sdk"
cd ../property_manager/ui
npm link localwebservices-sdk
