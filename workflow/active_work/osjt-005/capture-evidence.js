const { chromium } = require('playwright');
const fs = require('fs');

(async () => {
  const browser = await chromium.launch({headless: true});
  const page = await browser.newPage();
  
  const response = await page.goto('http://127.0.0.1:18219/', {waitUntil: 'domcontentloaded'});
  console.log('[Capture] HTTP Status:', response.status());
  
  const screenshotPath = 'workflow/active_work/osjt-005/evidence/screenshots/osjt-005-home.png';
  await page.screenshot({path: screenshotPath, fullPage: true});
  console.log('[Capture] Saved to:', require('path').resolve(screenshotPath));
  
  const title = await page.title();
  console.log('[Capture] Page Title:', title || '(no title)');
  console.log('[Capture] Content length:', (await page.content()).length, 'bytes');
  
  await browser.close();
  console.log('[Capture] Evidence capture complete!');
})();
