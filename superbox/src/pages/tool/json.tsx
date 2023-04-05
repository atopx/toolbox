import React from "react";

export interface EditorPanelProps {
    editable?: boolean;
    language?: string;
    defaultValue: string;
    title: React.ReactNode;
    hasCopy?: boolean;
    hasPrettier?: boolean;
    id: string | number;
    onChange?: (value: string) => void;
    hasLoad?: boolean;
    hasClear?: boolean;
    settingElement?: (args: { toggle: () => void; open: boolean }) => JSX.Element;
    alertMessage?: React.ReactNode;
    topNotifications?: (args: {
        toggleSettings: () => void;
        isSettingsOpen: boolean;
    }) => React.ReactNode;
    previewElement?: (value: string) => React.ReactNode;
    acceptFiles?: string | string[];
    packageDetails?: {
        name: string;
        url: string;
    };
}

export default function Editor() {
    return (
        <div>
            waiting
        </div>
    );
}