#!/usr/bin/env node
import * as cdk from 'aws-cdk-lib';
import { SmallurlStack } from 'smallurl-stack';

const app = new cdk.App();
new SmallurlStack(app, 'SmallurlStack');
