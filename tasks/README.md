# Task Definitions

This directory contains task definition files in YAML format.

## Purpose

Tasks define automated workflows that the engine will execute. Each task file specifies:
- Task metadata (name, description)
- Steps to execute
- Actions and parameters
- Dependencies

## Example

See `/examples/example-task.yaml` for a sample task definition.

## Usage

Place task YAML files in this directory. The task manager will discover and load them at runtime.
