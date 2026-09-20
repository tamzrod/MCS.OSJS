const { test, expect } = require('@playwright/test');

test.describe('OSJT-005 Basic Browser Automation', () => {
  test.beforeAll(async ({ browser }) => {
    console.log('[OSJT-005] Browser launched');
  });

  test.afterAll(async ({ browser }) => {
    await browser.close();
    console.log('[OSJT-005] Browser closed');
  });

  test('navigate to OSJS shell and capture screenshot', async ({ page }) => {
    console.log('[Step 1] Navigating to osjt-005 test URL...');
    
    const response = await page.goto('http://127.0.0.1:18219/', { waitUntil: 'domcontentloaded', timeout: 30000 });
    expect(response.status()).toBe(200);
    
    const screenshotPath = require('path').join(process.cwd(), 'work/evidence/screenshots/osjt-005-home.png');
    await page.screenshot({ path: screenshotPath, fullPage: true });
    console.log('[Step 2] Screenshot saved to:', screenshotPath);
    
    const contentLength = (await page.content()).length;
    expect(contentLength).toBeGreaterThan(100);
    
    console.log('[OSJT-005] Navigation successful!');
  });
});
