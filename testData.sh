#!/bin/bash

curl -H "Content-Type: application/json" -X POST -d '{"name":"xyz","password":"xyz","name2":"name2","password2":"abcd1234","effectiveDate":"12-12-12 04:00"}' http://localhost:33000/dropOffKeys
