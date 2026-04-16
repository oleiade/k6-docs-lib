---
title: 'exportKey'
description: 'exportKey exports a key in an external, portable format.'
weight: 04
---

# exportKey

The `exportKey()` method takes a [CryptoKey](https://grafana.com/docs/k6/<K6_VERSION>/javascript-api/crypto/cryptokey) object as input and exports it in an external, portable format.

Note that for a key to be exportable, it must have been created with the `extractable` flag set to `true`.

## Usage

```
exportKey(format, key)
```

## Parameters

| Name     | Type                                                                                  | Description                                                                                                                                                                                          |
| :------- | :------------------------------------------------------------------------------------ | :--------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `format` | `string`                                                                              | Defines the data format in which the key should be exported. Depending on the algorithm and key type, the data format could vary. Currently supported formats are `raw`, `jwk`, `spki`, and `pkcs8`. |
| `key`    | [CryptoKey](https://grafana.com/docs/k6/<K6_VERSION>/javascript-api/crypto/cryptokey) | The [key](https://grafana.com/docs/k6/<K6_VERSION>/javascript-api/crypto/cryptokey) to export.                                                                                                       |

### Supported algorithms


| AES-CBC                                                                                        | AES-CTR                                                                                        | AES-GCM                                                                                        | AES-KW | ECDH                                                                                                          | ECDSA                                                                                         | HMAC                                                                                                    | RSA-OAEP                                                                                                          | RSASSA-PKCS1-v1_5                                                                                                 | RSA-PSS                                                                                                           |
| :--------------------------------------------------------------------------------------------- | :--------------------------------------------------------------------------------------------- | :--------------------------------------------------------------------------------------------- | :----- | :------------------------------------------------------------------------------------------------------------ | :-------------------------------------------------------------------------------------------- | :------------------------------------------------------------------------------------------------------ | :---------------------------------------------------------------------------------------------------------------- | :---------------------------------------------------------------------------------------------------------------- | :---------------------------------------------------------------------------------------------------------------- |
| ✅ [AesCbcParams](https://grafana.com/docs/k6/<K6_VERSION>/javascript-api/crypto/aescbcparams) | ✅ [AesCtrParams](https://grafana.com/docs/k6/<K6_VERSION>/javascript-api/crypto/aesctrparams) | ✅ [AesGcmParams](https://grafana.com/docs/k6/<K6_VERSION>/javascript-api/crypto/aesgcmparams) | ❌     | ✅ [EcdhKeyDeriveParams](https://grafana.com/docs/k6/<K6_VERSION>/javascript-api/crypto/ecdhkeyderiveparams/) | ✅ [EcdsaParams](https://grafana.com/docs/k6/<K6_VERSION>/javascript-api/crypto/ecdsaparams/) | ✅ [HmacKeyGenParams](https://grafana.com/docs/k6/<K6_VERSION>/javascript-api/crypto/hmackeygenparams/) | ✅ [RsaHashedImportParams](https://grafana.com/docs/k6/<K6_VERSION>/javascript-api/crypto/rsahashedimportparams/) | ✅ [RsaHashedImportParams](https://grafana.com/docs/k6/<K6_VERSION>/javascript-api/crypto/rsahashedimportparams/) | ✅ [RsaHashedImportParams](https://grafana.com/docs/k6/<K6_VERSION>/javascript-api/crypto/rsahashedimportparams/) |


### Supported formats


- `ECDH` and `ECDSA` algorithms have support for `pkcs8`, `spki`, `raw` and `jwk` formats.
- `RSA-OAEP`, `RSASSA-PKCS1-v1_5` and `RSA-PSS` algorithms have support for `pkcs8`, `spki` and `jwk` formats.
- `AES-*` and `HMAC` algorithms have currently support for `raw` and `jwk` formats.


## Return Value

A `Promise` that resolves to a new `ArrayBuffer` or an [JsonWebKey](https://grafana.com/docs/k6/<K6_VERSION>/javascript-api/crypto/jsonwebkey) object/dictionary containing the key.

## Throws

| Type                 | Description                                         |
| :------------------- | :-------------------------------------------------- |
| `InvalidAccessError` | Raised when trying to export a non-extractable key. |
| `NotSupportedError`  | Raised when trying to export in an unknown format.  |
| `TypeError`          | Raised when trying to use an invalid format.        |

## Example

{{< code >}}

```javascript
export default async function () {
  /**
   * Generate a symmetric key using the AES-CBC algorithm.
   */
  const generatedKey = await crypto.subtle.generateKey(
    {
      name: 'AES-CBC',
      length: '256',
    },
    true,
    ['encrypt', 'decrypt']
  );

  /**
   * Export the key in raw format.
   */
  const exportedKey = await crypto.subtle.exportKey('raw', generatedKey);

  /**
   * Reimport the key in raw format to verify its integrity.
   */
  const importedKey = await crypto.subtle.importKey('raw', exportedKey, 'AES-CBC', true, [
    'encrypt',
    'decrypt',
  ]);

  console.log(JSON.stringify(importedKey));
}
```

{{< /code >}}
