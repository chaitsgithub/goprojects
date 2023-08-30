import { ComponentFixture, TestBed } from '@angular/core/testing';

import { JobstatsComponent } from './jobstats.component';

describe('JobstatsComponent', () => {
  let component: JobstatsComponent;
  let fixture: ComponentFixture<JobstatsComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ JobstatsComponent ]
    })
    .compileComponents();

    fixture = TestBed.createComponent(JobstatsComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
