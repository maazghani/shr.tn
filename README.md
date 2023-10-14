# shr.tn: Detailed Specification

## Introduction

The `shr.tn` application serves as an intuitive and streamlined gateway to GitHub repositories and potentially other social media platforms. By navigating to `shr.tn/`, users can enter minimally unique character sequences to find and share GitHub repositories effortlessly. For instance, navigating to `shr.tn/opnai/gpt` could redirect to `OpenAI/gpt-2`, and `shr.tn/tens` might lead to `tensorflow/tensorflow`. The aim is to simplify the sharing and discovery of repositories by replacing long URLs with short, easily shareable links.

Additionally, while the primary focus is GitHub, the foundational logic for fuzzy matching and redirection could readily be extended to other platforms like Facebook, LinkedIn, Instagram, Twitter, and YouTube.

---

## Summary

Built with OpenFaaS and Go, the `shr.tn` application consists of a custom router (`RequestRouter`) and two serverless functions (`ExecuteFuzzySearch`, `ManageOutboundRouting`). The development will strictly adhere to Test-Driven Development (TDD) principles, especially for the Go functions.

---

## Architecture Overview

### Components

1. **User Interface**: Custom domain `shr.tn/` for user input.
2. **RequestRouter**: HTTP router for URL parsing and directive handling.
3. **OpenFaaS Gateway**: Routes to the appropriate serverless function.
4. **Serverless Functions**: `ExecuteFuzzySearch` and `ManageOutboundRouting`.
5. **Backend**: GitHub API.

---

## Features and Operations

### RequestRouter

- **Input**: Raw URL.
- **Output**: Parsed URL and HTTP header directive (`X-Function-Invocation`).

### ExecuteFuzzySearch

- **Input**: Parsed strings (GitHub username or username/repo).
- **Output**: Best-matching GitHub username and/or repository.

### ManageOutboundRouting

- **Input**: Best-matching GitHub username and/or repository.
- **Output**: HTTP redirect to GitHub repository.

---

## Development Milestones

1. **Multi-Stage Dockerfile Creation**: 
    - Stage 1: Go build environment.
    - Stage 2: Barebones runtime environment.
  
2. **Initial Function Prototyping using TDD**: 
    - Create skeleton Go functions (`ExecuteFuzzySearch`, `ManageOutboundRouting`) following Test-Driven Development.

3. **Dockerization of Functions**: 
    - Containerize the Go functions using the multi-stage Dockerfile.

4. **OpenFaaS Environment Setup**: 
    - Deploy OpenFaaS and configure the Gateway.

5. **RequestRouter Implementation**: 
    - Develop and test the `RequestRouter`.

6. **Function Logic Implementation using TDD**: 
    - Complete core logic of `ExecuteFuzzySearch` and `ManageOutboundRouting` following Test-Driven Development.

7. **GitHub API Integration**: 
    - Implement API calls for fuzzy search functionality.

8. **Testing**: 
    - Unit tests, integration tests, and load tests.

9. **Deployment and Scaling**: 
    - Deploy to production and implement scaling strategies.

10. **Documentation**: 
    - Code comments, README, and end-user documentation.

---

## Test Plans

1. **Unit Tests**: Utilize Go's testing framework for function logic.
2. **Integration Tests**: Validate GitHub API and OpenFaaS integration.
3. **Load Tests**: Examine performance under simulated high-traffic scenarios.

