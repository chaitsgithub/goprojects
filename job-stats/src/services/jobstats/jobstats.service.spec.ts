import { TestBed } from '@angular/core/testing';

import { JobstatsService } from './jobstats.service';

describe('JobstatsService', () => {
  let service: JobstatsService;

  beforeEach(() => {
    TestBed.configureTestingModule({});
    service = TestBed.inject(JobstatsService);
  });

  it('should be created', () => {
    expect(service).toBeTruthy();
  });
});
