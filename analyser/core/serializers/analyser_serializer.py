from rest_framework import serializers


class _WorkloadSerializer(serializers.Serializer):

    interactions = serializers.IntegerField(min_value=100)
    intervalBetweenBatchesInMilliseconds = serializers.IntegerField(min_value=0)
    batchSize = serializers.IntegerField(min_value=1)

    # TODO: add custom field validation, add raise for uninmplemented methods

    def update(self, instance, validated_data):
        pass

    def create(self, validated_data):
        pass


class AnalyzerSerializer(serializers.Serializer):

    label = serializers.CharField()
    workload = _WorkloadSerializer()
    mean = serializers.FloatField()
    standardDeviation = serializers.FloatField()
    failures = serializers.IntegerField()
    rttsInMicroseconds = serializers.ListField(
       child=serializers.IntegerField()
    )

    def update(self, instance, validated_data):
        pass

    def create(self, validated_data):
        pass


