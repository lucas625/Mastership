from django.urls import path

from core.views import AnalyzerView
from core.views.ping_calculator_view import PingCalculatorView

urlpatterns = [
    path('analyze', AnalyzerView.as_view()),
    path('ping', PingCalculatorView.as_view())
]
