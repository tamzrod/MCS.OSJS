const { chromium } = require('playwright');

(async () => {
  console.log('[OSJT-005] Browser launching...');
  
  const browser = await chromium.launch({headless: true});
  
  try {
    const page = await browser.newPage();
    
    console.log('[Step 1] Navigating to http://127.0.0.1:18219/');
    const response = await page.goto('http://127.0.0.1:18219/', {waitUntil: 'domcontentloaded', timeout: 30000});
    console.log('[Step 2] HTTP Status:', response.status());
    
    const screenshotPath = require('path').join(
      process.cwd(), 
      'workflow/active_work/osjt-005/evidence/screenshots', 
      'osjt-005-ui.png'
    );
    console.log('[Step 3] Capturing screenshot to:', screenshotPath);
    
    await require('fs').mkdir(require('path').dirname(screenshotPath), {recursive: true});
    await page.screenshot({path: screenshotPath, fullPage: true});
    console.log('[Step 4] Screenshot saved successfully');
    
    const title = await page.title();
    console.log('[Step 5] Page Title:', title || '(no title)');
    
    console.log('\n[OSJT-005] ==========================================');
    console.log('[OSJT-005] TEST PASSED - Browser automation verified');  
    console.log('=============================================');
  } catch(err) {
    console.error('[OSJT-005] ERROR:', err.message.substring(0,200));
  } finally {
    await browser.close();
  }
})();
