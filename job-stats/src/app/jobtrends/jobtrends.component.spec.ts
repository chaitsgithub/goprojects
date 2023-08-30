import { ComponentFixture, TestBed } from '@angular/core/testing';

import { JobtrendsComponent } from './jobtrends.component';

describe('JobtrendsComponent', () => {
  let component: JobtrendsComponent;
  let fixture: ComponentFixture<JobtrendsComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ JobtrendsComponent ]
    })
    .compileComponents();

    fixture = TestBed.createComponent(JobtrendsComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
