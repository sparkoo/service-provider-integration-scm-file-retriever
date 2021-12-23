# service-provider-integration-scm-file-retriever
Package for downloading files from a source code management sites

### About

This repository contains a library for retrieving files from various source management systems using a repositories and a file paths as the primary form of input.

The main idea is to allow its users to download files from a different SCM providers without the necessity of knowing their APIs and/or download endpoints,
as well as take care about the authentication.

### Usage



### URL and path formats
Repository URLs may or may not contain `.git` suffixes. Paths are usual `/a/b/filename` format. Optional `ref` may 
contain commit id, tag or branch name.

### Supported SCM providers