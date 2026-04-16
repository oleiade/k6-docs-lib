---
title: 'getByTestId(testId)'
description: 'Browser module: locator.getByTestId(testId) method'
---


# getByTestId(testId)

Returns a locator for elements with the specified test ID attribute. This method is designed for robust test automation by locating elements using dedicated test identifiers that are independent of the visual appearance or content changes. Currently it can only work with the `data-testid` attribute.

| Parameter | Type             | Default | Description                                                                                                                                            |
| --------- | ---------------- | ------- | ------------------------------------------------------------------------------------------------------------------------------------------------------ |
| `testId`  | string \| RegExp | -       | Required. The test ID value to search for. Searches for the `data-testid` attribute. Can be a string for exact match or a RegExp for pattern matching. |

## Returns

| Type                                                                                   | Description                                                                                     |
| -------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------- |
| [Locator](https://grafana.com/docs/k6/<K6_VERSION>/javascript-api/k6-browser/locator/) | A locator object that can be used to interact with the elements matching the specified test ID. |

## Examples

### Basic test ID usage

Locate and interact with elements using test IDs:

```javascript
import { browser } from 'k6/browser';

export const options = {
  scenarios: {
    browser: {
      executor: 'shared-iterations',
      options: {
        browser: {
          type: 'chromium',
        },
      },
    },
  },
};

export default async function () {
  const page = await browser.newPage();

  try {
    await page.setContent(`
      <input type="text" data-testid="username">
      <input type="text" data-testid="email">
      <button data-testid="submit-button">Submit</button>
    `);

    const locator = page.locator(':root');
    await locator.getByTestId('username').fill('FirstLast');
    await locator.getByTestId('email').fill('firstlast@example.com');
    await locator.getByTestId('submit-button').click();
  } finally {
    await page.close();
  }
}
```


## Best practices

1. **Stable identifiers**: Use meaningful, stable test IDs that won't change with refactoring or content updates.
1. **Hierarchical naming**: Use consistent naming conventions like `user-profile-edit-btn`.
1. **Avoid duplicates**: Ensure test IDs are unique within the page to prevent ambiguity.
1. **Strategic placement**: Add test IDs to key interactive elements and components that are frequently tested.
1. **Team coordination**: Establish test ID conventions with your development team to ensure consistency.

## Related

- [locator.getByRole()](https://grafana.com/docs/k6/<K6_VERSION>/javascript-api/k6-browser/locator/getbyrole/) - Locate by ARIA role
- [locator.getByAltText()](https://grafana.com/docs/k6/<K6_VERSION>/javascript-api/k6-browser/locator/getbyalttext/) - Locate by alt text
- [locator.getByLabel()](https://grafana.com/docs/k6/<K6_VERSION>/javascript-api/k6-browser/locator/getbylabel/) - Locate by form labels
- [locator.getByPlaceholder()](https://grafana.com/docs/k6/<K6_VERSION>/javascript-api/k6-browser/locator/getbyplaceholder/) - Locate by placeholder text
- [locator.getByText()](https://grafana.com/docs/k6/<K6_VERSION>/javascript-api/k6-browser/locator/getbytext/) - Locate by text content
- [locator.getByTitle()](https://grafana.com/docs/k6/<K6_VERSION>/javascript-api/k6-browser/locator/getbytitle/) - Locate by title attribute
