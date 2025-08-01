import { ApplicationConfig, provideZoneChangeDetection } from '@angular/core';
import { provideRouter } from '@angular/router';
import { provideHttpClient, withInterceptorsFromDi } from '@angular/common/http';
import { registerLocaleData } from '@angular/common';
import localeTh from '@angular/common/locales/th';
import { LOCALE_ID } from '@angular/core';

import { routes } from './app.routes';

registerLocaleData(localeTh);

export const appConfig: ApplicationConfig = {
  providers: [
    provideZoneChangeDetection({ eventCoalescing: true }),
    provideRouter(routes),
    provideHttpClient(withInterceptorsFromDi()),  //แบบใหม่ของ Angular 19.2.15 (my)
    {provide: LOCALE_ID, useValue: 'th-TH'}
  ],
};
