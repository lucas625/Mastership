from rest_framework import serializers

from core.data import Analyzer
from core.serializers.workload_serializer import WorkloadSerializer


class AnalyzerSerializer(serializers.Serializer):

    label = serializers.CharField()
    workload = WorkloadSerializer()
    mean = serializers.FloatField()
    standardDeviation = serializers.FloatField()
    failures = serializers.IntegerField()
    rttsInMilliseconds = serializers.ListField(
        child=serializers.FloatField(),
        min_length=1,
    )

    def validate(self, data):
        if data['failures'] > data['workload']['interactions']:
            raise serializers.ValidationError({'failures': 'it can\'t be greater than interactions.'})
        if len(data['rttsInMilliseconds']) != data['workload']['interactions']:
            raise serializers.ValidationError({'rttsInMilliseconds': 'it must be equal to interactions.'})
        return data

    def update(self, instance, validated_data):
        raise NotImplementedError("This method is not implemented for this class")

    def create(self, validated_data):
        workload_serializer = WorkloadSerializer(data=validated_data['workload'])
        workload_serializer.is_valid()
        return Analyzer(
            validated_data['label'],
            workload_serializer.save(),
            validated_data['mean'],
            validated_data['standardDeviation'],
            validated_data['failures'],
            validated_data['rttsInMilliseconds'],
        )
