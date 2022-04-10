# azcopy

This images packages [azcopy](https://github.com/Azure/azure-storage-azcopy).

## Running the image

### Azure Blob Storage

```bash
target=$(printf "https://%s.blob.core.windows.net/%s?%s" \
  "$AZURE_STORAGE_ACCOUNT" \
  "$AZURE_STORAGE_CONTAINER" \
  "$AZURE_STORAGE_SAS_TOKEN")

docker run -ti --rm -e target -v "$PWD/src:/src" \
    bluebrown/azcopy sync src/ "$target"
```
